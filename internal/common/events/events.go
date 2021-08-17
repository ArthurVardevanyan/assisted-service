// Code generated by from events definition; DO NOT EDIT.
package events

import (
    "context"
    "fmt"
    "strings"
    "time"
    "github.com/openshift/assisted-service/internal/events"

    "github.com/go-openapi/strfmt"
)

//
// Event cancel_install_failed_start
//
type CancelInstallFailedStartEvent struct {
    ClusterId strfmt.UUID
}

var CancelInstallFailedStartEventName string = "cancel_install_failed_start"

func NewCancelInstallFailedStartEvent(
    clusterId strfmt.UUID,
) *CancelInstallFailedStartEvent {
    return &CancelInstallFailedStartEvent{
        ClusterId: clusterId,
    }
}

func SendCancelInstallFailedStartEvent(
    ctx context.Context,
    eventsHandler events.Sender,
    clusterId strfmt.UUID,) {
    ev := NewCancelInstallFailedStartEvent(
        clusterId,
    )
    eventsHandler.SendClusterEvent(ctx, ev)
}

func SendCancelInstallFailedStartEventAtTime(
    ctx context.Context,
    eventsHandler events.Sender,
    clusterId strfmt.UUID,
    eventTime time.Time) {
    ev := NewCancelInstallFailedStartEvent(
        clusterId,
    )
    eventsHandler.SendClusterEventAtTime(ctx, ev, eventTime)
}

func (e *CancelInstallFailedStartEvent) GetName() string {
    return "cancel_install_failed_start"
}

func (e *CancelInstallFailedStartEvent) GetSeverity() string {
    return "error"
}

func (e *CancelInstallFailedStartEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

func (e *CancelInstallFailedStartEvent) format(message *string) string {
    r := strings.NewReplacer(
        "{cluster_id}", fmt.Sprint(e.ClusterId),
    )
    return r.Replace(*message)
}

func (e *CancelInstallFailedStartEvent) FormatMessage() string {
    s := "Failed to cancel installation: error starting DB transaction"
    return e.format(&s)
}

//
// Event cancel_install_failed_commit
//
type CancelInstallFailedCommitEvent struct {
    ClusterId strfmt.UUID
}

var CancelInstallFailedCommitEventName string = "cancel_install_failed_commit"

func NewCancelInstallFailedCommitEvent(
    clusterId strfmt.UUID,
) *CancelInstallFailedCommitEvent {
    return &CancelInstallFailedCommitEvent{
        ClusterId: clusterId,
    }
}

func SendCancelInstallFailedCommitEvent(
    ctx context.Context,
    eventsHandler events.Sender,
    clusterId strfmt.UUID,) {
    ev := NewCancelInstallFailedCommitEvent(
        clusterId,
    )
    eventsHandler.SendClusterEvent(ctx, ev)
}

func SendCancelInstallFailedCommitEventAtTime(
    ctx context.Context,
    eventsHandler events.Sender,
    clusterId strfmt.UUID,
    eventTime time.Time) {
    ev := NewCancelInstallFailedCommitEvent(
        clusterId,
    )
    eventsHandler.SendClusterEventAtTime(ctx, ev, eventTime)
}

func (e *CancelInstallFailedCommitEvent) GetName() string {
    return "cancel_install_failed_commit"
}

func (e *CancelInstallFailedCommitEvent) GetSeverity() string {
    return "error"
}

func (e *CancelInstallFailedCommitEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}

func (e *CancelInstallFailedCommitEvent) format(message *string) string {
    r := strings.NewReplacer(
        "{cluster_id}", fmt.Sprint(e.ClusterId),
    )
    return r.Replace(*message)
}

func (e *CancelInstallFailedCommitEvent) FormatMessage() string {
    s := "Failed to cancel installation: error committing DB transaction"
    return e.format(&s)
}

//
// Event host_registration_failed_setting_properties
//
type HostRegistrationFailedSettingPropertiesEvent struct {
    ClusterId strfmt.UUID
    HostId strfmt.UUID
}

var HostRegistrationFailedSettingPropertiesEventName string = "host_registration_failed_setting_properties"

func NewHostRegistrationFailedSettingPropertiesEvent(
    clusterId strfmt.UUID,
    hostId strfmt.UUID,
) *HostRegistrationFailedSettingPropertiesEvent {
    return &HostRegistrationFailedSettingPropertiesEvent{
        ClusterId: clusterId,
        HostId: hostId,
    }
}

func SendHostRegistrationFailedSettingPropertiesEvent(
    ctx context.Context,
    eventsHandler events.Sender,
    clusterId strfmt.UUID,
    hostId strfmt.UUID,) {
    ev := NewHostRegistrationFailedSettingPropertiesEvent(
        clusterId,
        hostId,
    )
    eventsHandler.SendHostEvent(ctx, ev)
}

func SendHostRegistrationFailedSettingPropertiesEventAtTime(
    ctx context.Context,
    eventsHandler events.Sender,
    clusterId strfmt.UUID,
    hostId strfmt.UUID,
    eventTime time.Time) {
    ev := NewHostRegistrationFailedSettingPropertiesEvent(
        clusterId,
        hostId,
    )
    eventsHandler.SendHostEventAtTime(ctx, ev, eventTime)
}

func (e *HostRegistrationFailedSettingPropertiesEvent) GetName() string {
    return "host_registration_failed_setting_properties"
}

func (e *HostRegistrationFailedSettingPropertiesEvent) GetSeverity() string {
    return "error"
}

func (e *HostRegistrationFailedSettingPropertiesEvent) GetClusterId() *strfmt.UUID {
    return &e.ClusterId
}
func (e *HostRegistrationFailedSettingPropertiesEvent) GetHostId() *strfmt.UUID {
    return &e.HostId
}

func (e *HostRegistrationFailedSettingPropertiesEvent) format(message *string) string {
    r := strings.NewReplacer(
        "{cluster_id}", fmt.Sprint(e.ClusterId),
        "{host_id}", fmt.Sprint(e.HostId),
    )
    return r.Replace(*message)
}

func (e *HostRegistrationFailedSettingPropertiesEvent) FormatMessage() string {
    s := "Failed to register host: error setting host properties"
    return e.format(&s)
}

