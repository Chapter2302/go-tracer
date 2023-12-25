package tracer

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jounng23/go-tracer/config"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgrpc"
	"go.elastic.co/apm/transport"
	"google.golang.org/grpc"
)

func (t apmTracer) getUnaryInterceptors() []grpc.UnaryServerInterceptor {
	if t.tracer == nil {
		return nil
	}
	serverOpts := apmgrpc.WithTracer(t.tracer)
	return []grpc.UnaryServerInterceptor{
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
		apmgrpc.NewUnaryServerInterceptor(serverOpts),
	}
}

func (t *apmTracer) initGrpcTracer(svcName, svcVer string) (err error) {
	if t == nil {
		return errors.New("Empty tracer")
	}

	var aTracer *apm.Tracer

	apm.DefaultTracer.Close()
	os.Setenv("ELASTIC_APM_SERVER_URL", config.ApmConfig().Url)
	os.Setenv("ELASTIC_APM_ACTIVE", config.ApmConfig().Active)
	os.Setenv("ELASTIC_APM_METRICS_INTERVAL", config.ApmConfig().MetricsInterval)
	os.Setenv("ELASTIC_APM_VERIFY_SERVER_CERT", config.ApmConfig().VerifyCert)
	os.Setenv("ELASTIC_APM_TRANSACTION_SAMPLE_RATE", config.ApmConfig().SampleRate)
	os.Setenv("ELASTIC_APM_SPAN_FRAMES_MIN_DURATION", config.ApmConfig().SpanMinDuration)

	_, _ = transport.InitDefault()
	aTracer, err = apm.NewTracer(svcName, svcVer)
	if err == nil && aTracer == nil {
		err = errors.New("Empty APM tracer")
	}

	if err != nil {
		return
	}

	t.tracer = aTracer
	return
}

func (tracer apmTracer) startSpanFromCtx(ctx context.Context, spanName string) (ISpan, context.Context) {
	span, ctx := apm.StartSpan(ctx, spanName, "")
	return &apmSpan{span}, ctx
}

func (a apmSpan) End() {
	if a.span == nil {
		return
	}
	a.span.End()
}

func (a *apmSpan) SetTag(key string, val interface{}) {
	if a.span == nil {
		return
	}

	if key == "type" {
		if str, ok := val.(StrSpanType); ok {
			a.span.Type = fmt.Sprint(str)
		}
		return
	}
	a.span.Context.SetTag(key, fmt.Sprint(val))
	return
}
