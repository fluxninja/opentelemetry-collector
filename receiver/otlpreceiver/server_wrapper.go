package otlpreceiver

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
)

var _ ptraceotlp.Server = (*TraceServerWrapper)(nil)

// TraceServerWrapper is a thin wrapper around ptraceotlp.Server. It can be registered
// in grpc.Server and the underlying ptraceotlp.Server can be set even after the
// grpc.Server is started.
type TraceServerWrapper struct {
	server ptraceotlp.Server
}

func (s *TraceServerWrapper) Export(ctx context.Context, r ptraceotlp.Request) (ptraceotlp.Response, error) {
	if s.server == nil {
		return ptraceotlp.NewResponse(), errors.New("handler not initialized")
	}
	return s.server.Export(ctx, r)
}

var _ pmetricotlp.Server = (*MetricServerWrapper)(nil)

// MetricServerWrapper is a thin wrapper around pmetricotlp.Server. It can be registered
// in grpc.Server and the underlying pmetricotlp.Server can be set even after the
// grpc.Server is started.
type MetricServerWrapper struct {
	server pmetricotlp.Server
}

func (s *MetricServerWrapper) Export(ctx context.Context, r pmetricotlp.Request) (pmetricotlp.Response, error) {
	if s.server == nil {
		return pmetricotlp.NewResponse(), errors.New("handler not initialized")
	}
	return s.server.Export(ctx, r)
}

var _ plogotlp.Server = (*LogServerWrapper)(nil)

// LogServerWrapper is a thin wrapper around plogotlp.Server. It can be registered
// in grpc.Server and the underlying plogotlp.Server can be set even after the
// grpc.Server is started.
type LogServerWrapper struct {
	server plogotlp.Server
}

func (s *LogServerWrapper) Export(ctx context.Context, r plogotlp.Request) (plogotlp.Response, error) {
	if s.server == nil {
		return plogotlp.NewResponse(), errors.New("handler not initialized")
	}
	return s.server.Export(ctx, r)
}
