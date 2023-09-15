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

func SetupContextTracerUsecase(ctx context.Context) (context.Context, *TraceSpan) {
	/*
		ctx = logv2.InitFuncContext(ctx)
		ctx = tracer.AddNewrelicAttribute(ctx)
		ctx, tSpan := tracer.StartSpanFromContext(ctx)
	*/

	return ctx, nil
}
