// Package performance provides the Chrome DevTools Protocol
// commands, types, and events for the Performance domain.
//
// Generated by the cdproto-gen command.
package performance

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/cdvelop/vanify/pkg/cdproto/cdp"
)

// DisableParams disable collecting and reporting metrics.
type DisableParams struct{}

// Disable disable collecting and reporting metrics.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Performance#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Performance.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enable collecting and reporting metrics.
type EnableParams struct {
	TimeDomain EnableTimeDomain `json:"timeDomain,omitempty"` // Time domain to use for collecting and reporting duration metrics.
}

// Enable enable collecting and reporting metrics.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Performance#method-enable
//
// parameters:
func Enable() *EnableParams {
	return &EnableParams{}
}

// WithTimeDomain time domain to use for collecting and reporting duration
// metrics.
func (p EnableParams) WithTimeDomain(timeDomain EnableTimeDomain) *EnableParams {
	p.TimeDomain = timeDomain
	return &p
}

// Do executes Performance.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, p, nil)
}

// GetMetricsParams retrieve current values of run-time metrics.
type GetMetricsParams struct{}

// GetMetrics retrieve current values of run-time metrics.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Performance#method-getMetrics
func GetMetrics() *GetMetricsParams {
	return &GetMetricsParams{}
}

// GetMetricsReturns return values.
type GetMetricsReturns struct {
	Metrics []*Metric `json:"metrics,omitempty"` // Current values for run-time metrics.
}

// Do executes Performance.getMetrics against the provided context.
//
// returns:
//
//	metrics - Current values for run-time metrics.
func (p *GetMetricsParams) Do(ctx context.Context) (metrics []*Metric, err error) {
	// execute
	var res GetMetricsReturns
	err = cdp.Execute(ctx, CommandGetMetrics, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Metrics, nil
}

// Command names.
const (
	CommandDisable    = "Performance.disable"
	CommandEnable     = "Performance.enable"
	CommandGetMetrics = "Performance.getMetrics"
)
