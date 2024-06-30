// Package extensions provides the Chrome DevTools Protocol
// commands, types, and events for the Extensions domain.
//
// Defines commands and events for browser extensions. Available if the
// client is connected using the --remote-debugging-pipe flag and the
// --enable-unsafe-extension-debugging flag is set.
//
// Generated by the cdproto-gen command.
package extensions

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
)

// LoadUnpackedParams installs an unpacked extension from the filesystem
// similar to --load-extension CLI flags. Returns extension ID once the
// extension has been installed.
type LoadUnpackedParams struct {
	Path string `json:"path"` // Absolute file path.
}

// LoadUnpacked installs an unpacked extension from the filesystem similar to
// --load-extension CLI flags. Returns extension ID once the extension has been
// installed.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Extensions#method-loadUnpacked
//
// parameters:
//
//	path - Absolute file path.
func LoadUnpacked(path string) *LoadUnpackedParams {
	return &LoadUnpackedParams{
		Path: path,
	}
}

// LoadUnpackedReturns return values.
type LoadUnpackedReturns struct {
	ID string `json:"id,omitempty"` // Extension id.
}

// Do executes Extensions.loadUnpacked against the provided context.
//
// returns:
//
//	id - Extension id.
func (p *LoadUnpackedParams) Do(ctx context.Context) (id string, err error) {
	// execute
	var res LoadUnpackedReturns
	err = cdp.Execute(ctx, CommandLoadUnpacked, p, &res)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}

// Command names.
const (
	CommandLoadUnpacked = "Extensions.loadUnpacked"
)
