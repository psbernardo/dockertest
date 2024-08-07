package tracing

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type SpanErrorType string

const (
	SPAN_INTERNAL_ERROR         SpanErrorType = "Internal Error"
	SPAN_DATABASE_ERROR         SpanErrorType = "Database Error"
	SPAN_RECORD_NOT_FOUND_ERROR SpanErrorType = "Record Not Found Error"
	SPAN_VALIDATION_ERROR       SpanErrorType = "Validation Error"
	SPAN_PAYLOAD_PARSE_ERROR    SpanErrorType = "Payload Parse Error"
)

type TracerSpan struct {
	trace.Span
	layer    TracerLayer
	spanName string
}

func (t *TracerSpan) SetAttributes(kv ...attribute.KeyValue) {
	t.Span.SetAttributes(kv...)
}

func (t *TracerSpan) AddValues(values map[string]interface{}) {
	for key, value := range values {
		t.Span.SetAttributes(attribute.Key(key).String(fmt.Sprint(value)))
	}
}

func (t *TracerSpan) End(errs ...error) {
	for _, err := range errs {
		t.reportError(err)
	}

	t.Span.End()
}

func (t *TracerSpan) reportError(err error) {
	if err == nil {
		t.Span.SetStatus(codes.Ok, fmt.Sprintf("%s Successfully executed...", t.spanName))
		return
	}

	t.SetStatus(codes.Error, err.Error())
	t.RecordError(err, trace.WithStackTrace(true))

}
