// Package preload provides the Chrome DevTools Protocol
// commands, types, and events for the Preload domain.
//
// Generated by the cdproto-gen command.
package preload

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
)

// EnableParams [no description].
type EnableParams struct{}

// Enable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Preload.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// DisableParams [no description].
type DisableParams struct{}

// Disable [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Preload#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Preload.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// Command names.
const (
	CommandEnable  = "Preload.enable"
	CommandDisable = "Preload.disable"
)
