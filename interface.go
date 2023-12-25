package tracer

import (
	"context"

	"google.golang.org/grpc"
)

type ITracer interface {
	initGrpcTracer(svcName, svcVer string) (err error)
	startSpanFromCtx(ctx context.Context, spanName string) (ISpan, context.Context)
	getUnaryInterceptors() []grpc.UnaryServerInterceptor
}

type ISpan interface {
	SetTag(key string, val interface{})
	End()
}
