// Package deviceaccess provides the Chrome DevTools Protocol
// commands, types, and events for the DeviceAccess domain.
//
// Generated by the cdproto-gen command.
package deviceaccess

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
)

// EnableParams enable events in this domain.
type EnableParams struct{}

// Enable enable events in this domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DeviceAccess#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes DeviceAccess.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// DisableParams disable events in this domain.
type DisableParams struct{}

// Disable disable events in this domain.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DeviceAccess#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes DeviceAccess.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// SelectPromptParams select a device in response to a
// DeviceAccess.deviceRequestPrompted event.
type SelectPromptParams struct {
	ID       RequestID `json:"id"`
	DeviceID DeviceID  `json:"deviceId"`
}

// SelectPrompt select a device in response to a
// DeviceAccess.deviceRequestPrompted event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DeviceAccess#method-selectPrompt
//
// parameters:
//
//	id
//	deviceID
func SelectPrompt(id RequestID, deviceID DeviceID) *SelectPromptParams {
	return &SelectPromptParams{
		ID:       id,
		DeviceID: deviceID,
	}
}

// Do executes DeviceAccess.selectPrompt against the provided context.
func (p *SelectPromptParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSelectPrompt, p, nil)
}

// CancelPromptParams cancel a prompt in response to a
// DeviceAccess.deviceRequestPrompted event.
type CancelPromptParams struct {
	ID RequestID `json:"id"`
}

// CancelPrompt cancel a prompt in response to a
// DeviceAccess.deviceRequestPrompted event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DeviceAccess#method-cancelPrompt
//
// parameters:
//
//	id
func CancelPrompt(id RequestID) *CancelPromptParams {
	return &CancelPromptParams{
		ID: id,
	}
}

// Do executes DeviceAccess.cancelPrompt against the provided context.
func (p *CancelPromptParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandCancelPrompt, p, nil)
}

// Command names.
const (
	CommandEnable       = "DeviceAccess.enable"
	CommandDisable      = "DeviceAccess.disable"
	CommandSelectPrompt = "DeviceAccess.selectPrompt"
	CommandCancelPrompt = "DeviceAccess.cancelPrompt"
)
