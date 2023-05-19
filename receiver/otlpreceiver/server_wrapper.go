package otlpreceiver

import (
	"context"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	switchable_handler switchableHandler[ptraceotlp.ExportRequest, ptraceotlp.ExportResponse]
}

func (s *TraceServerWrapper) Export(
	ctx context.Context,
	r ptraceotlp.ExportRequest,
) (ptraceotlp.ExportResponse, error) {
	return s.switchable_handler.Export(ctx, r)
}

var _ pmetricotlp.GRPCServer = (*MetricServerWrapper)(nil)

// MetricServerWrapper is a thin wrapper around pmetricotlp.GRPCServer. It can be registered
// in grpc.GRPCServer and the underlying pmetricotlp.GRPCServer can be set even after the
// grpc.GRPCServer is started.
type MetricServerWrapper struct {
	pmetricotlp.UnimplementedGRPCServer
	switchable_handler switchableHandler[pmetricotlp.ExportRequest, pmetricotlp.ExportResponse]
}

func (s *MetricServerWrapper) Export(
	ctx context.Context,
	r pmetricotlp.ExportRequest,
) (pmetricotlp.ExportResponse, error) {
	return s.switchable_handler.Export(ctx, r)
}

var _ plogotlp.GRPCServer = (*LogServerWrapper)(nil)

// LogServerWrapper is a thin wrapper around plogotlp.GRPCServer. It can be registered
// in grpc.GRPCServer and the underlying plogotlp.GRPCServer can be set even after the
// grpc.GRPCServer is started.
type LogServerWrapper struct {
	plogotlp.UnimplementedGRPCServer
	switchable_handler switchableHandler[plogotlp.ExportRequest, plogotlp.ExportResponse]
}

func (s *LogServerWrapper) Export(
	ctx context.Context,
	r plogotlp.ExportRequest,
) (plogotlp.ExportResponse, error) {
	return s.switchable_handler.Export(ctx, r)
}

// switchableHandler is a thin generic wrapper around
// {trace,metrics,logs}.Receiver, which can be set / unset in runtime.
// Note: It doesn't implement p{trace,metric,log}otlp.GRPCSever directly,
// because of missing *.UnimplementedGRPCServer.
type switchableHandler[Req any, Resp any] struct {
	// nil when collector is loading/reloading
	exporter atomic.Pointer[shutdownableExporter[Req, Resp]]
}

// Shutdown unsets currently active handler and marks server as temporarily
// unavailable, waiting for all active requests to finish.
//
// After Shutdown is called, Set/SetUnimplemented can be called another time.
func (sw *switchableHandler[Req, Resp]) Shutdown() {
	exporter := sw.exporter.Load()
	if exporter != nil {
		// Mark collector as reloading
		sw.exporter.Store(nil)
		// Finish ongoing requests
		exporter.room.Close()
	}
}

// Set sets the exporter as currently active handler
func (sw *switchableHandler[Req, Resp]) Set(exporter exporter[Req, Resp]) {
	sw.exporter.Store(&shutdownableExporter[Req, Resp]{exporter: exporter})
}

// SetUnimplemented sets an Unimplemented-returning handler as active.
func (sw *switchableHandler[Req, Resp]) SetUnimplemented() {
	// Wrapping unimplementedExporter in shutdownableExporter is unnecessary as
	// there's nothing to shut down there, but it's simpler this way.
	sw.exporter.Store(&shutdownableExporter[Req, Resp]{
		exporter: unimplementedExporter[Req, Resp]{},
	})
}

func (sw *switchableHandler[Req, Resp]) Export(ctx context.Context, req Req) (Resp, error) {
	exporter := sw.exporter.Load()
	if exporter == nil {
		var resp Resp
		return resp, status.Error(codes.Unavailable, "collector reloading")
		// Note: This means that we respond with Unavailable in case when no
		// otlp receiver is configured and thus otlpReceiver is not being
		// started at all. We don't expect such scenario though.
	}
	return exporter.Export(ctx, req)
}

// shutdownableExporter wraps an exporter allowing graceful shutdown (waiting
// for inflight requests and blocking subsequent ones)
type shutdownableExporter[Req any, Resp any] struct {
	room     Room
	exporter exporter[Req, Resp]
}

func (e *shutdownableExporter[Req, Resp]) Export(ctx context.Context, req Req) (Resp, error) {
	if ok := e.room.TryEnter(); !ok {
		var noResp Resp
		return noResp, status.Error(codes.Unavailable, "collector reloading")
	}
	defer e.room.Leave()
	return e.exporter.Export(ctx, req)
}

func (e *shutdownableExporter[Req, Resp]) Shutdown() { e.room.Close() }

// exporter is a helper interface for all of ptraceotlp, pmetricotlp, plogsotlp.GRPCServer
type exporter[Req any, Resp any] interface {
	Export(ctx context.Context, r Req) (Resp, error)
}

type unimplementedExporter[Req any, Resp any] struct{}

func (_ unimplementedExporter[Req, Resp]) Export(ctx context.Context, req Req) (Resp, error) {
	var resp Resp
	return resp, status.Error(codes.Unimplemented, "receiver not enabled")
}
