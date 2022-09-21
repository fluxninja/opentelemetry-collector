// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpreceiver // import "go.opentelemetry.io/collector/receiver/otlpreceiver"

import (
	"go.opentelemetry.io/collector/component"
)

// Config defines configuration for OTLP receiver.
type Config struct {
	traceServerWrapper  *TraceServerWrapper
	metricServerWrapper *MetricServerWrapper
	logServerWrapper    *LogServerWrapper
}

var _ component.Config = (*Config)(nil)
