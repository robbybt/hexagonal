package context

import (
	"context"
)

type TraceSpan struct {
	span        *tracing.Span
	transaction *tracing.Transaction
}

func SetupContextTracerUsecase(ctx context.Context) (context.Context, *TraceSpan) {
	/*
		ctx = logv2.InitFuncContext(ctx)
		ctx = tracer.AddNewrelicAttribute(ctx)
		ctx, tSpan := tracer.StartSpanFromContext(ctx)
	*/

	return ctx, nil
}
