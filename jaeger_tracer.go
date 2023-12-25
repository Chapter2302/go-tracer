package tracer

import (
	"context"
	"errors"

	"github.com/jounng23/go-tracer/config"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	jaegerClient "github.com/uber/jaeger-client-go"
	jaegerClientCfg "github.com/uber/jaeger-client-go/config"
)

func (t jaegerTracer) getUnaryInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_middleware.ChainUnaryServer(
			otgrpc.OpenTracingServerInterceptor(t.tracer), // Jaeger tracer interceptor
		),
	}
}

func (t *jaegerTracer) initGrpcTracer(svcName, svcVer string) (err error) {
	if t == nil {
		return errors.New("Empty tracer")
	}

	var jTracer opentracing.Tracer

	cfg, err := config.GetJaegerCfgFromEnv(svcName)
	if err != nil {
		return
	}

	jTracer, _, err = cfg.NewTracer(jaegerClientCfg.Logger(jaegerClient.StdLogger))
	if err != nil {
		return
	}

	opentracing.SetGlobalTracer(jTracer)
	t.tracer = jTracer
	return
}

func (tracer jaegerTracer) startSpanFromCtx(ctx context.Context, spanName string) (ISpan, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, spanName)
	return &jaegerSpan{span: span}, ctx
}

func (j jaegerSpan) End() {
	j.span.Finish()
}

func (j *jaegerSpan) SetTag(key string, val interface{}) {
	j.span.SetTag(key, val)
	return
}
