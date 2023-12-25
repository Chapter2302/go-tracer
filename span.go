package tracer

import (
	"context"
)

func StartSpanFromCtx(ctx context.Context, spanName string,
	tags ...spanTag) (ISpan, context.Context) {
	if currTracer == nil {
		return nil, nil
	}
	span, ctx := currTracer.startSpanFromCtx(ctx, spanName)
	for _, tag := range tags {
		span.SetTag(tag.Key, tag.Value)
	}
	return span, ctx
}

func WithType(val StrSpanType) spanTag {
	return spanTag{
		Key:   "type",
		Value: val,
	}
}
