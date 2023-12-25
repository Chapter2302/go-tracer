package tracer

import (
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm"
)

type apmTracer struct {
	tracer *apm.Tracer
}

type jaegerTracer struct {
	tracer opentracing.Tracer
}

type spanTag struct {
	Key   string
	Value interface{}
}

type spanType func() StrSpanType

type jaegerSpan struct {
	span opentracing.Span
}

type apmSpan struct {
	span *apm.Span
}
