package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc/metadata"
)

const traceIDKey = "x-trace-id"

// Tracing is a middleware that integrates OpenTracing with the Echo framework.
// It creates a span for each incoming request and propagates the trace context
// using gRPC metadata and response headers. The trace ID is set in the response
// headers with the key "x-trace-id".
//
// If the request processing results in an error, the span is marked with the error
// and the error message is logged in the trace. Otherwise, a success response tag
// is added to the span. The middleware ensures that the tracing context is properly
// propagated by updating the request context.
func Tracing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		span, ctx := opentracing.StartSpanFromContext(ctx, c.Request().Method)
		defer span.Finish()

		if spanContext, ok := span.Context().(jaeger.SpanContext); ok {
			traceID := spanContext.TraceID().String()

			c.Response().Header().Set(traceIDKey, traceID)

			ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(traceIDKey, traceID))
		}

		c.SetRequest(c.Request().WithContext(ctx))

		err := next(c)

		if err != nil {
			ext.Error.Set(span, true)
			span.SetTag("error", err.Error())
		} else {
			span.SetTag("response", "api gateway success response")
		}

		return err
	}
}
