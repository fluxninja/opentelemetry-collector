// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "go.opentelemetry.io/collector/exporter/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/exporter/internal/otlptext"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/fluxninja/aperture/pkg/log"
)

type loggingExporter struct {
	logsMarshaler    plog.Marshaler
	metricsMarshaler pmetric.Marshaler
	tracesMarshaler  ptrace.Marshaler
}

func (s *loggingExporter) pushTraces(_ context.Context, td ptrace.Traces) error {
	log.Trace().Int("#spans", td.SpanCount()).
		Int("#resource spans", td.ResourceSpans().Len()).
		Msg("TracesExporter")

	buf, err := s.tracesMarshaler.MarshalTraces(td)
	if err != nil {
		return err
	}
	log.Trace().Msg(string(buf))
	return nil
}

func (s *loggingExporter) pushMetrics(_ context.Context, md pmetric.Metrics) error {
	log.Trace().Int("#metrics", md.MetricCount()).
		Int("#resource metrics", md.ResourceMetrics().Len()).
		Int("#data points", md.DataPointCount()).
		Msg("MetricsExporter")

	buf, err := s.metricsMarshaler.MarshalMetrics(md)
	if err != nil {
		return err
	}
	log.Trace().Msg(string(buf))
	return nil
}

func (s *loggingExporter) pushLogs(_ context.Context, ld plog.Logs) error {
	log.Trace().Int("#logs", ld.LogRecordCount()).
		Int("#resource logs", ld.ResourceLogs().Len()).
		Msg("LogsExporter")

	buf, err := s.logsMarshaler.MarshalLogs(ld)
	if err != nil {
		return err
	}
	log.Trace().Msg(string(buf))
	return nil
}

func newLoggingExporter() *loggingExporter {
	return &loggingExporter{
		logsMarshaler:    otlptext.NewTextLogsMarshaler(),
		metricsMarshaler: otlptext.NewTextMetricsMarshaler(),
		tracesMarshaler:  otlptext.NewTextTracesMarshaler(),
	}
}
