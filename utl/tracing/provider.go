package tracing

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTraceProvider(ctx context.Context, conf Config) (*trace.TracerProvider, error) {
	exporter, err := newExporter(ctx, conf)
	if err != nil {
		return nil, errors.New("unable to create trace provider")
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(newResource(conf)),
	), nil
}
