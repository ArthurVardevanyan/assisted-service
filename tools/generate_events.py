#!/usr/bin/env python3

import yaml
import logging
import argparse
import jinja2
from collections import OrderedDict


EVENTS_DEFINITION = 'x-events'

# NOQA
EVENT_TEMPLATE = jinja2.Template('''// Code generated by from events definition; DO NOT EDIT.
package {{package_name}}

import (
    "context"
    "fmt"
    "strings"
    "time"
    "github.com/openshift/assisted-service/internal/events"
{% for import in extra_imports %}
    {{import}}
{% endfor -%}
)

{% for event in generator -%}
//
// Event {{event['name']}}
//
{% set eventName = event.event_class() -%}
{% set eventBaseName = event.package_private(eventName) + "Base" -%}
{% set baseEvent = event.base_event(event) -%}
type {{eventName}} struct {
{%- for p, t in event.properties.items() %}
    {{event.pascal_case(p)}} {{event.go_type(t)}}
{%- endfor %}
}

var {{eventName}}Name string = "{{event['name']}}"

func New{{eventName}}(
{%- for p, t in event.properties.items() %}
    {{event.camel_case(p)}} {{event.go_type(t)}},
{%- endfor %}
) *{{eventName}} {
    return &{{eventName}}{
{%- for p, t in event.properties.items() %}
        {{event.pascal_case(p)}}: {{event.camel_case(p)}},
{%- endfor %}
    }
}

func Send{{eventName}}(
    ctx context.Context,
    eventsHandler events.Sender,
{%- for p, t in event.properties.items() %}
    {{event.camel_case(p)}} {{event.go_type(t)}},
{%- endfor %}) {
    ev := New{{eventName}}(
{%- for p, t in event.properties.items() %}
        {{event.camel_case(p)}},
{%- endfor %}
    )
    eventsHandler.Send{{event.pascal_case(event.type)}}Event(ctx, ev)
}

func Send{{eventName}}AtTime(
    ctx context.Context,
    eventsHandler events.Sender,
{%- for p, t in event.properties.items() %}
    {{event.camel_case(p)}} {{event.go_type(t)}},
{%- endfor %}
    eventTime time.Time) {
    ev := New{{eventName}}(
{%- for p, t in event.properties.items() %}
        {{event.camel_case(p)}},
{%- endfor %}
    )
    eventsHandler.Send{{event.pascal_case(event.type)}}EventAtTime(ctx, ev, eventTime)
}

func (e *{{eventName}}) GetName() string {
    return "{{event['name']}}"
}

func (e *{{eventName}}) GetSeverity() string {
    {% if event.properties.severity -%}     return e.Severity
    {%- else -%}      return "{{event.event_severity()}}"
    {%- endif %}
}

func (e *{{eventName}}) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

{%- if event.type != "cluster" %}
func (e *{{eventName}}) GetHostId() *strfmt.UUID {
    return &e.HostId
}
{%- endif %}

func (e *{{eventName}}) format(message *string) string {
    r := strings.NewReplacer(
{%- for p, t in event['properties'].items() %}
        "{{'{'}}{{p}}{{'}'}}", fmt.Sprint(e.{{event.pascal_case(p)}}),
{%- endfor %}
    )
    return r.Replace(*message)
}

func (e *{{eventName}}) FormatMessage() string {
    s := "{{event.format}}"
    return e.format(&s)
}

{% endfor -%}
''')


class EventGenerator(object):
    VALID_PROPS_TYPES = {"UUID":dict(go_type="strfmt.UUID", go_import="github.com/go-openapi/strfmt"),
                         "integer":dict(go_type="int", go_import=None),
                         "int64":dict(go_type="int64", go_import=None),
                         "string":dict(go_type="string", go_import=None),
                         "bool":dict(go_type="bool", go_import=None)}
    def __init__(self, event):
        self.event = event

    def __getattr__(self, attr):
        #import ipdb; ipdb.set_trace()
        if attr in self.__dict__:
            return getattr(self, attr)
        else:
            return getattr(self.event, attr)

    def ctor_name(self):
        return "New" + self.event_class()

    def ctor_args(self):
        args = []
        for arg, _ in self.event.properties.iteritems():
            args.append(EventGenerator.camel_case(arg) + " " + EventGenerator.pascal_case(arg))
        return ", ".join(args)

    def event_class(self):
        return EventGenerator.pascal_case(self.event.name) + "Event"

    def event_type(self):
        return "Type" + EventGenerator.pascal_case(self.event.name)

    def event_severity(self):
        return self.event.severity.lower()

    @staticmethod
    def base_event(event):
        if event.type == "cluster":
            return "ClusterBaseEvent"
        else:
            return "HostBaseEvent"

    @staticmethod
    def cap(s):
        if len(s) == 0:
            return s
        return s[0].upper() + s[1:]

    @staticmethod
    def package_private(s):
        if len(s) == 0:
            return s
        return s[0].lower() + s[1:]

    @staticmethod
    def camel_case(value):
        parts = value.split('_')
        return parts[0] + ''.join(EventGenerator.cap(part) for part in parts[1:])

    @staticmethod
    def pascal_case(value):
        return ''.join(EventGenerator.cap(part) for part in value.split('_'))

    @staticmethod
    def go_type(var_type):
        return EventGenerator.VALID_PROPS_TYPES[var_type]['go_type']

    @staticmethod
    def type_to_string(var_type, var_name):
        if var_type == 'string':
            return var_name
        else:
            return 'fmt.Sprint({})'.format(var_name)


class EventDef:
    def __init__(self, name, format, event_type, severity, properties):
        self.name = name
        self.format = format
        self.type = event_type
        self.severity = severity
        self.properties = properties

def validate_property_types_and_return_extra_imports(e):
    extra_imports = set()
    for t in e['properties'].values():
        if t not in EventGenerator.VALID_PROPS_TYPES:
            raise Exception("Invalid property type ('{}') for '{}' in {}".format(t, p, e['name']))
        else:
            if EventGenerator.VALID_PROPS_TYPES[t]['go_import']:
                extra_imports.add('"{}"'.format(EventGenerator.VALID_PROPS_TYPES[t]['go_import']))
    return extra_imports

def validate_event(e):
    REQUIRED_CLUSTER_PROPERTIES = ["cluster_id"]
    REQUIRED_HOST_PROPERTIES = ["cluster_id", "host_id"]
    INVALID_CLUSTER_PROPERTIES = ["host_id"]
    INVALID_HOST_PROPERTIES = []
    VALID_SEVERITY_VALUES = ["info", "warning", "error", "critical"]

    required_props = REQUIRED_HOST_PROPERTIES
    if e['event_type'] == "cluster":
        required_props = REQUIRED_CLUSTER_PROPERTIES
    for p in required_props:
        if p not in e['properties']:
            raise Exception("Missing '{}' in properties of {}".format(p, e['name']))

    invalid_props = INVALID_HOST_PROPERTIES
    if e['event_type'] == "cluster":
        invalid_props = INVALID_CLUSTER_PROPERTIES
    for p in invalid_props:
        if p in e['properties']:
            raise Exception("Invalid '{}' in properties of {}".format(p, e['name']))

    valid_severities = VALID_SEVERITY_VALUES
    if e['severity'] not in valid_severities:
        raise Exception("Invalid '{}' as severity of {}".format(e['severity'], e['name']))


def parse(yaml_path):
    def dict_ctor(loader, node):
        return OrderedDict(loader.construct_pairs(node))

    # Need to use ordered dict with yaml.Loader so that the order of definitions will not change
    # each time we generate
    # This is specifically important for the order of the event ctor arguments
    yaml.Loader.add_constructor(yaml.resolver.BaseResolver.DEFAULT_MAPPING_TAG, dict_ctor)
    with open(yaml_path, "r") as spec_file:
        spec = yaml.load(spec_file, Loader=yaml.Loader)

    if spec is None:
        raise ValueError("Invalid input file")

    events_spec = spec.get(EVENTS_DEFINITION)
    if events_spec is None:
        logging.error("Cannot find events in %s", yaml_path)
        raise ValueError(
            "Invalid input file, missing %r section" % EVENTS_DEFINITION
        )

    events = []
    extra_imports = set()
    event_names = set()
    for e in events_spec:
        if e['name'] not in event_names:
            event_names.add(e['name'])
        else:
            raise Exception("Duplicate event name {} found".format(e['name']))
        validate_event(e)
        extra_imports |= validate_property_types_and_return_extra_imports(e)
        ne = EventDef(name=e['name'], format=e['message'],
                      event_type=e['event_type'], severity=e['severity'],
                      properties=e['properties'])
        events.append(ne)

    return events, extra_imports


def main():
    """Tool entry point - read parameters from users and start tool."""
    parser = argparse.ArgumentParser(
        description='Generate events definitions based on specifications.'
    )
    parser.add_argument('source', type=str, help='yaml source path')
    parser.add_argument('dest', type=str, help='destination file path')
    parser.add_argument('-p', '--package-name', default="events", help='go package name')
    args = parser.parse_args()

    events, extra_imports = parse(yaml_path=args.source)
    generated_code = EVENT_TEMPLATE.render(generator=(EventGenerator(e) for e in events),
                                           extra_imports=extra_imports,
                                           package_name=args.package_name)

    with open(args.dest, "w+") as fout:
        fout.write(generated_code)


if __name__ == '__main__':
    main()
