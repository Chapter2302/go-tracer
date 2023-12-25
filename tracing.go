package tracer

import (
	"google.golang.org/grpc"
)

var currTracer ITracer

func CurrentTracer() ITracer {
	return currTracer
}

func InitGrpcTracer(svcName, svcVer string,
	eTracer ETracer) (interceptors []grpc.UnaryServerInterceptor, err error) {
	switch eTracer {
	case JAEGER:
		currTracer = &jaegerTracer{}
	case APM:
		currTracer = &apmTracer{}
	}

	err = currTracer.initGrpcTracer(svcName, svcVer)
	if err != nil {
		return
	}
	return currTracer.getUnaryInterceptors(), nil
}
