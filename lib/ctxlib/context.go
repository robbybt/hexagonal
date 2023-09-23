package ctxlib

import (
	"context"
)

type TraceSpan struct {
	a string
}

func (t *TraceSpan) Finish() {
	return
}

// SetupContextTracerUsecase will inject context with tracer new relic for tracing
func SetupContextTracerUsecase(ctx context.Context) (context.Context, *TraceSpan) {
	/*
		ctx = logv2.InitFuncContext(ctx)
		ctx = tracer.AddNewrelicAttribute(ctx)
		ctx, tSpan := tracer.StartSpanFromContext(ctx)
	*/

	return ctx, nil
}
