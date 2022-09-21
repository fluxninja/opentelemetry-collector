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
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/logs"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/metrics"
	"go.opentelemetry.io/collector/receiver/otlpreceiver/internal/trace"
	"google.golang.org/grpc"
)

// otlpReceiver is the type that exposes Trace and Metrics reception.
type otlpReceiver struct {
	cfg        *Config
	serverGRPC *grpc.Server
	httpMux    *http.ServeMux
	serverHTTP *http.Server

	traceReceiver   *trace.Receiver
	metricsReceiver *metrics.Receiver
	logReceiver     *logs.Receiver
	shutdownWG      sync.WaitGroup

	settings receiver.CreateSettings
}

// newOtlpReceiver just creates the OpenTelemetry receiver services. It is the caller's
// responsibility to invoke the respective Start*Reception methods as well
// as the various Stop*Reception methods to end it.
func newOtlpReceiver(cfg *Config, settings receiver.CreateSettings) *otlpReceiver {
	return &otlpReceiver{
		cfg:      cfg,
		settings: settings,
	}
}

// Shutdown is a method to turn off receiving.
func (r *otlpReceiver) Shutdown(ctx context.Context) error {
	return nil
}

// Start runs the trace receiver on the gRPC server. Currently
// it also enables the metrics receiver too.
func (r *otlpReceiver) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (r *otlpReceiver) registerTraceConsumer(tc consumer.Traces) error {
	if tc == nil {
		return component.ErrNilNextConsumer
	}
	var err error
	r.traceReceiver, err = trace.New(tc, r.settings)
	if err != nil {
		return err
	}
	r.cfg.traceServerWrapper.server = r.traceReceiver
	return nil
}

func (r *otlpReceiver) registerMetricsConsumer(mc consumer.Metrics) error {
	if mc == nil {
		return component.ErrNilNextConsumer
	}
	var err error
	r.metricsReceiver, err = metrics.New(mc, r.settings)
	if err != nil {
		return err
	}
	r.cfg.metricServerWrapper.server = r.metricsReceiver
	return nil
}

func (r *otlpReceiver) registerLogsConsumer(lc consumer.Logs) error {
	if lc == nil {
		return component.ErrNilNextConsumer
	}
	var err error
	r.logReceiver, err = logs.New(lc, r.settings)
	if err != nil {
		return err
	}
	r.cfg.logServerWrapper.server = r.logReceiver
	return nil
}
