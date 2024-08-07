package tracing

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type TracerLayer string

const (
	TRACE_LAYER_USECASE    TracerLayer = "Usecase"
	TRACE_LAYER_HANDLER    TracerLayer = "Handler"
	TRACE_LAYER_REPOSITORY TracerLayer = "Repository"
)

func NewTraceHanlder() Tracer {
	return newTrace(TRACE_LAYER_HANDLER)
}

func NewTraceUseCase() Tracer {
	return newTrace(TRACE_LAYER_USECASE)
}

func NewTraceRepository() Tracer {
	return newTrace(TRACE_LAYER_REPOSITORY)
}

func (t TracerLayer) String() string {
	return string(t)
}

type Tracer struct {
	layer TracerLayer
}

func newTrace(layer TracerLayer) Tracer {
	return Tracer{
		layer: layer,
	}
}

func (t Tracer) StartSpan(ctx context.Context) (context.Context, TracerSpan) {
	fnName := getFnName(2)
	fnName = t.Fname(fnName)
	spanContext, span := otel.Tracer(fnName).Start(ctx, fnName)
	span.SetAttributes(
		attribute.Key("Layer").String(t.layer.String()),
	)

	return spanContext, TracerSpan{
		layer:    TracerLayer(t.layer),
		Span:     span,
		spanName: fnName,
	}
}

func (t Tracer) StartSpanWithAttributes(ctx context.Context, attributes map[string]interface{}) (context.Context, TracerSpan, error) {
	fnName := getFnName(2)
	fnName = t.Fname(fnName)
	spanContext, span := otel.Tracer(fnName).Start(ctx, fnName)
	span.SetAttributes(
		attribute.Key("Layer").String(t.layer.String()),
	)

	for key, value := range attributes {
		span.SetAttributes(attribute.Key(key).String(fmt.Sprint(value)))
	}
	return spanContext, TracerSpan{
		layer:    TracerLayer(t.layer),
		Span:     span,
		spanName: fnName,
	}, nil
}

func (t Tracer) Fname(fnName string) string {
	return fmt.Sprintf("%s.%s", t.layer, fnName)
}

func getFnName(skip int) string {
	counter, _, _, _ := runtime.Caller(skip)
	fnSlict := strings.Split(runtime.FuncForPC(counter).Name(), ".")
	return fnSlict[len(fnSlict)-1]
}
