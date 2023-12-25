package config

import (
	"errors"

	jaegerClientCfg "github.com/uber/jaeger-client-go/config"
)

const (
	DEFAULT_JAEGER_SAMPLER_TYPE  = "const"
	DEFAULT_JAEGER_SAMPLER_PARAM = 1
)

var DefaultJaegerSamplerCfg = jaegerClientCfg.SamplerConfig{
	Type:  DEFAULT_JAEGER_SAMPLER_TYPE,
	Param: DEFAULT_JAEGER_SAMPLER_PARAM,
}

func GetJaegerCfgFromEnv(svcName string) (cfg *jaegerClientCfg.Configuration, err error) {
	cfg, err = jaegerClientCfg.FromEnv()
	if err != nil {
		return
	}

	if cfg.ServiceName == "" {
		if svcName == "" {
			err = errors.New("empty service name")
			return
		}
		cfg.ServiceName = svcName
	}

	if cfg.Sampler != nil && cfg.Sampler.Type == "" {
		cfg.Sampler = &DefaultJaegerSamplerCfg
	}

	if cfg.Reporter != nil && !cfg.Reporter.LogSpans {
		cfg.Reporter.LogSpans = true
	}
	return
}
