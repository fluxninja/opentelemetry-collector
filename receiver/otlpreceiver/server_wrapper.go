package otlpreceiver

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
)

var _ ptraceotlp.GRPCServer = (*TraceServerWrapper)(nil)

// TraceServerWrapper is a thin wrapper around ptraceotlp.GRPCServer. It can be registered
// in grpc.GRPCServer and the underlying ptraceotlp.GRPCServer can be set even after the
// grpc.GRPCServer is started.
type TraceServerWrapper struct {
	ptraceotlp.UnimplementedGRPCServer
	server ptraceotlp.GRPCServer
}

func (s *TraceServerWrapper) Export(ctx context.Context, r ptraceotlp.ExportRequest) (ptraceotlp.ExportResponse, error) {
	if s.server == nil {
		return ptraceotlp.NewExportResponse(), errors.New("handler not initialized")
	}
	return s.server.Export(ctx, r)
}

var _ pmetricotlp.GRPCServer = (*MetricServerWrapper)(nil)

// MetricServerWrapper is a thin wrapper around pmetricotlp.GRPCServer. It can be registered
// in grpc.GRPCServer and the underlying pmetricotlp.GRPCServer can be set even after the
// grpc.GRPCServer is started.
type MetricServerWrapper struct {
	pmetricotlp.UnimplementedGRPCServer
	server pmetricotlp.GRPCServer
}

func (s *MetricServerWrapper) Export(ctx context.Context, r pmetricotlp.ExportRequest) (pmetricotlp.ExportResponse, error) {
	if s.server == nil {
		return pmetricotlp.NewExportResponse(), errors.New("handler not initialized")
	}
	return s.server.Export(ctx, r)
}

var _ plogotlp.GRPCServer = (*LogServerWrapper)(nil)

// LogServerWrapper is a thin wrapper around plogotlp.GRPCServer. It can be registered
// in grpc.GRPCServer and the underlying plogotlp.GRPCServer can be set even after the
// grpc.GRPCServer is started.
type LogServerWrapper struct {
	plogotlp.UnimplementedGRPCServer
	server plogotlp.GRPCServer
}

func (s *LogServerWrapper) Export(ctx context.Context, r plogotlp.ExportRequest) (plogotlp.ExportResponse, error) {
	if s.server == nil {
		return plogotlp.NewExportResponse(), errors.New("handler not initialized")
	}
	return s.server.Export(ctx, r)
}
