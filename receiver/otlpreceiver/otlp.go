// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otlpreceiver // import "go.opentelemetry.io/collector/receiver/otlpreceiver"

import (
	"context"
	"sync"

	"google.golang.org/grpc"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/obsreport"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/logs"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/metrics"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/trace"
)

// otlpReceiver is the type that exposes Trace and Metrics reception.
type otlpReceiver struct {
	cfg        *Config
	serverGRPC *grpc.Server

	tracesReceiver  *trace.Receiver
	metricsReceiver *metrics.Receiver
	logsReceiver    *logs.Receiver

	obsrepGRPC *obsreport.Receiver

	settings receiver.CreateSettings
}

// newOtlpReceiver just creates the OpenTelemetry receiver services. It is the caller's
// responsibility to invoke the respective Start*Reception methods as well
// as the various Stop*Reception methods to end it.
func newOtlpReceiver(cfg *Config, set receiver.CreateSettings) (*otlpReceiver, error) {
	r := &otlpReceiver{
		cfg:      cfg,
		settings: set,
	}
	var err error
	r.obsrepGRPC, err = obsreport.NewReceiver(obsreport.ReceiverSettings{
		ReceiverID:             set.ID,
		Transport:              "grpc",
		ReceiverCreateSettings: set,
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
