package interceptors

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const traceIDKey = "x-trace-id"

// ServerTracing is a gRPC server interceptor that integrates with OpenTracing for distributed tracing.
// It starts a new span for each incoming gRPC request, attaches the trace ID to the metadata, and
// sends it as part of the response headers.
// The method also captures errors and response information, annotating the span accordingly.
func ServerTracing(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// создаем отметку
	span, ctx := opentracing.StartSpanFromContext(ctx, info.FullMethod)
	defer span.Finish()

	// кастуем к типу SpanContext и если ok - делаем что надо
	spanContext, ok := span.Context().(jaeger.SpanContext)
	if ok {
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(traceIDKey, spanContext.TraceID().String()))
		header := metadata.New(map[string]string{traceIDKey: spanContext.TraceID().String()})
		err := grpc.SendHeader(ctx, header)
		if err != nil {
			return nil, err
		}
	}

	res, err := handler(ctx, req)

	if err != nil {
		ext.Error.Set(span, true) // рисует красный кружок ( УДОБНО )
		span.SetTag("error", err.Error())
	} else {
		span.SetTag("response", res)
	}

	return res, err
}
