package tracer

type StrSpanType string

const (
	HANDLER  StrSpanType = "handler"
	CACHE    StrSpanType = "cache"
	STORAGE  StrSpanType = "storage"
	EXTERNAL StrSpanType = "external"
)
