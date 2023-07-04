// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpreceiver // import "go.opentelemetry.io/collector/receiver/otlpreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/internal/sharedcomponent"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeStr = "otlp"
)

// NewFactory creates a new OTLP receiver factory.
func NewFactory(
	tsw *TraceServerWrapper,
	msw *MetricServerWrapper,
	lsw *LogServerWrapper,
) receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig(tsw, msw, lsw),
		receiver.WithTraces(createTraces, component.StabilityLevelStable),
		receiver.WithMetrics(createMetrics, component.StabilityLevelStable),
		receiver.WithLogs(createLog, component.StabilityLevelBeta))
}

// createDefaultConfig creates the default configuration for receiver.
func createDefaultConfig(
	tsw *TraceServerWrapper,
	msw *MetricServerWrapper,
	lsw *LogServerWrapper,
) component.CreateDefaultConfigFunc {
	return func() component.Config {
		return &Config{
			traceServerWrapper:  tsw,
			metricServerWrapper: msw,
			logServerWrapper:    lsw,
		}
	}
}

// createTraces creates a trace receiver based on provided config.
func createTraces(
	_ context.Context,
	set receiver.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	oCfg := cfg.(*Config)
	r, err := receivers.GetOrAdd(oCfg, func() (*otlpReceiver, error) {
		return newOtlpReceiver(oCfg, set)
	})
	if err != nil {
		return nil, err
	}

	if err = r.Unwrap().registerTraceConsumer(nextConsumer); err != nil {
		return nil, err
	}
	return r, nil
}

// createMetrics creates a metrics receiver based on provided config.
func createMetrics(
	_ context.Context,
	set receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	oCfg := cfg.(*Config)
	r, err := receivers.GetOrAdd(oCfg, func() (*otlpReceiver, error) {
		return newOtlpReceiver(oCfg, set)
	})
	if err != nil {
		return nil, err
	}

	if err = r.Unwrap().registerMetricsConsumer(consumer); err != nil {
		return nil, err
	}
	return r, nil
}

// createLog creates a log receiver based on provided config.
func createLog(
	_ context.Context,
	set receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	oCfg := cfg.(*Config)
	r, err := receivers.GetOrAdd(oCfg, func() (*otlpReceiver, error) {
		return newOtlpReceiver(oCfg, set)
	})
	if err != nil {
		return nil, err
	}

	if err = r.Unwrap().registerLogsConsumer(consumer); err != nil {
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
var receivers = sharedcomponent.NewSharedComponents[*Config, *otlpReceiver]()
