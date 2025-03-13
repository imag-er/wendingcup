package common

import (
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"os"
)

func InitTracing(serviceName string) provider.OtelProvider {
	var entryPoint string
	if os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") == "" {
		entryPoint = "localhost:4317"
	} else {
		entryPoint = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	}
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
		provider.WithExportEndpoint(entryPoint),
	)
	if p == nil {
		panic("failed to init tracing for service " + serviceName)
	}
	return p
}
