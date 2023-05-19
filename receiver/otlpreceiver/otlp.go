// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpreceiver // import "go.opentelemetry.io/collector/receiver/otlpreceiver"

import (
	"context"
	"sync"

	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/logs"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/metrics"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/trace"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// otlpReceiver is the type that exposes Trace and Metrics reception.
type otlpReceiver struct {
	cfg        *Config
	serverGRPC *grpc.Server

	tracesReceiver  *trace.Receiver
	metricsReceiver *metrics.Receiver
	logsReceiver    *logs.Receiver

	obsrepGRPC *receiverhelper.ObsReport
	obsrepHTTP *receiverhelper.ObsReport

	settings *receiver.CreateSettings
}

// newOtlpReceiver just creates the OpenTelemetry receiver services. It is the caller's
// responsibility to invoke the respective Start*Reception methods as well
// as the various Stop*Reception methods to end it.
func newOtlpReceiver(cfg *Config, set *receiver.CreateSettings) (*otlpReceiver, error) {
	r := &otlpReceiver{
		cfg:      cfg,
		settings: set,
	}
	var err error
	r.obsrepGRPC, err = receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             set.ID,
		Transport:              "grpc",
		ReceiverCreateSettings: *set,
	})
	if err != nil {
		return nil, err
	}
	r.obsrepHTTP, err = receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             set.ID,
		Transport:              "http",
		ReceiverCreateSettings: *set,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Shutdown is a method to turn off receiving.
func (r *otlpReceiver) Shutdown(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		r.cfg.traceServerWrapper.switchable_handler.Shutdown()
		wg.Done()
	}()

	go func() {
		r.cfg.metricServerWrapper.switchable_handler.Shutdown()
		wg.Done()
	}()

	go func() {
		r.cfg.logServerWrapper.switchable_handler.Shutdown()
		wg.Done()
	}()

	wg.Wait()
	return nil
}

// Start runs the trace receiver on the gRPC server. Currently
// it also enables the metrics receiver too.
func (r *otlpReceiver) Start(_ context.Context, _ component.Host) error {
	if r.tracesReceiver != nil {
		r.cfg.traceServerWrapper.switchable_handler.Set(r.tracesReceiver)
	} else {
		r.cfg.traceServerWrapper.switchable_handler.SetUnimplemented()
	}

	if r.metricsReceiver != nil {
		r.cfg.metricServerWrapper.switchable_handler.Set(r.metricsReceiver)
	} else {
		r.cfg.metricServerWrapper.switchable_handler.SetUnimplemented()
	}

	if r.logsReceiver != nil {
		r.cfg.logServerWrapper.switchable_handler.Set(r.logsReceiver)
	} else {
		r.cfg.logServerWrapper.switchable_handler.SetUnimplemented()
	}

	return nil
}

func (r *otlpReceiver) registerTraceConsumer(tc consumer.Traces) error {
	if tc == nil {
		return component.ErrNilNextConsumer
	}
	r.tracesReceiver = trace.New(tc, r.obsrepGRPC)
	return nil
}

func (r *otlpReceiver) registerMetricsConsumer(mc consumer.Metrics) error {
	if mc == nil {
		return component.ErrNilNextConsumer
	}
	r.metricsReceiver = metrics.New(mc, r.obsrepGRPC)
	return nil
}

func (r *otlpReceiver) registerLogsConsumer(lc consumer.Logs) error {
	if lc == nil {
		return component.ErrNilNextConsumer
	}
	r.logsReceiver = logs.New(lc, r.obsrepGRPC)
	return nil
}
