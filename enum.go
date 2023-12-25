package tracer

type ETracer int32

const (
	JAEGER ETracer = iota + 1
	APM
)
