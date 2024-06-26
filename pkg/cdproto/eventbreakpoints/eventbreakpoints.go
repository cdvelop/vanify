// Package eventbreakpoints provides the Chrome DevTools Protocol
// commands, types, and events for the EventBreakpoints domain.
//
// EventBreakpoints permits setting JavaScript breakpoints on operations and
// events occurring in native code invoked from JavaScript. Once breakpoint is
// hit, it is reported through Debugger domain, similarly to regular breakpoints
// being hit.
//
// Generated by the cdproto-gen command.
package eventbreakpoints

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
)

// SetInstrumentationBreakpointParams sets breakpoint on particular native
// event.
type SetInstrumentationBreakpointParams struct {
	EventName string `json:"eventName"` // Instrumentation name to stop on.
}

// SetInstrumentationBreakpoint sets breakpoint on particular native event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints#method-setInstrumentationBreakpoint
//
// parameters:
//
//	eventName - Instrumentation name to stop on.
func SetInstrumentationBreakpoint(eventName string) *SetInstrumentationBreakpointParams {
	return &SetInstrumentationBreakpointParams{
		EventName: eventName,
	}
}

// Do executes EventBreakpoints.setInstrumentationBreakpoint against the provided context.
func (p *SetInstrumentationBreakpointParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetInstrumentationBreakpoint, p, nil)
}

// RemoveInstrumentationBreakpointParams removes breakpoint on particular
// native event.
type RemoveInstrumentationBreakpointParams struct {
	EventName string `json:"eventName"` // Instrumentation name to stop on.
}

// RemoveInstrumentationBreakpoint removes breakpoint on particular native
// event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints#method-removeInstrumentationBreakpoint
//
// parameters:
//
//	eventName - Instrumentation name to stop on.
func RemoveInstrumentationBreakpoint(eventName string) *RemoveInstrumentationBreakpointParams {
	return &RemoveInstrumentationBreakpointParams{
		EventName: eventName,
	}
}

// Do executes EventBreakpoints.removeInstrumentationBreakpoint against the provided context.
func (p *RemoveInstrumentationBreakpointParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandRemoveInstrumentationBreakpoint, p, nil)
}

// DisableParams removes all breakpoints.
type DisableParams struct{}

// Disable removes all breakpoints.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/EventBreakpoints#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes EventBreakpoints.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// Command names.
const (
	CommandSetInstrumentationBreakpoint    = "EventBreakpoints.setInstrumentationBreakpoint"
	CommandRemoveInstrumentationBreakpoint = "EventBreakpoints.removeInstrumentationBreakpoint"
	CommandDisable                         = "EventBreakpoints.disable"
)
