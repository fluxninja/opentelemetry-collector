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

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/internal/sharedcomponent"
)

const (
	typeStr = "otlp"
)

// NewFactory creates a new OTLP receiver factory.
func NewFactory(
	tsw *TraceServerWrapper,
	msw *MetricServerWrapper,
	lsw *LogServerWrapper,
) component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig(tsw, msw, lsw),
		component.WithTracesReceiver(createTracesReceiver, component.StabilityLevelStable),
		component.WithMetricsReceiver(createMetricsReceiver, component.StabilityLevelStable),
		component.WithLogsReceiver(createLogReceiver, component.StabilityLevelBeta))
}

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig(
	tsw *TraceServerWrapper,
	msw *MetricServerWrapper,
	lsw *LogServerWrapper,
) component.ReceiverCreateDefaultConfigFunc {
	return func() config.Receiver {
		return &Config{
			ReceiverSettings:    config.NewReceiverSettings(config.NewComponentID(typeStr)),
			traceServerWrapper:  tsw,
			metricServerWrapper: msw,
			logServerWrapper:    lsw,
		}
	}
}

// createTracesReceiver creates a trace receiver based on provided config.
func createTracesReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	cfg config.Receiver,
	nextConsumer consumer.Traces,
) (component.TracesReceiver, error) {
	r := receivers.GetOrAdd(cfg, func() component.Component {
		return newOtlpReceiver(cfg.(*Config), set)
	})

	if err := r.Unwrap().(*otlpReceiver).registerTraceConsumer(nextConsumer); err != nil {
		return nil, err
	}
	return r, nil
}

// createMetricsReceiver creates a metrics receiver based on provided config.
func createMetricsReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	cfg config.Receiver,
	consumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	r := receivers.GetOrAdd(cfg, func() component.Component {
		return newOtlpReceiver(cfg.(*Config), set)
	})

	if err := r.Unwrap().(*otlpReceiver).registerMetricsConsumer(consumer); err != nil {
		return nil, err
	}
	return r, nil
}

// createLogReceiver creates a log receiver based on provided config.
func createLogReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	cfg config.Receiver,
	consumer consumer.Logs,
) (component.LogsReceiver, error) {
	r := receivers.GetOrAdd(cfg, func() component.Component {
		return newOtlpReceiver(cfg.(*Config), set)
	})

	if err := r.Unwrap().(*otlpReceiver).registerLogsConsumer(consumer); err != nil {
		return nil, err
	}
	return r, nil
}

// This is the map of already created OTLP receivers for particular configurations.
// We maintain this map because the Factory is asked trace and metric receivers separately
// when it gets CreateTracesReceiver() and CreateMetricsReceiver() but they must not
// create separate objects, they must use one otlpReceiver object per configuration.
// When the receiver is shutdown it should be removed from this map so the same configuration
// can be recreated successfully.
var receivers = sharedcomponent.NewSharedComponents()
