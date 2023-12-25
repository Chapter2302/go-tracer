package config

var (
	apmCfg ApmCfg
)

type ApmCfg struct {
	Url             string `envconfig:"ELASTIC_APM_SERVER_URL" default:"http://127.0.0.1:8200/"`
	VerifyCert      string `envconfig:"ELASTIC_APM_VERIFY_SERVER_CERT" default:"false"`
	Active          string `envconfig:"ELASTIC_APM_ACTIVE" default:"true"`
	SampleRate      string `envconfig:"ELASTIC_APM_TRANSACTION_SAMPLE_RATE" default:"1.0"`
	MetricsInterval string `envconfig:"ELASTIC_APM_METRICS_INTERVAL" default:"1s"`
	SpanMinDuration string `envconfig:"ELASTIC_APM_SPAN_FRAMES_MIN_DURATION" default:"0ms"`
}

func ApmConfig() ApmCfg {
	return apmCfg
}
