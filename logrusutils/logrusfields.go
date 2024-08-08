package logrusutils

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

var globalAPMService string

// LogrusFields returns logrus fields for tracing information
func LogrusFields(spanCtx trace.SpanContext) logrus.Fields {
	if !spanCtx.IsValid() || globalAPMService == "" {
		return logrus.Fields{}
	}
	return logrus.Fields{
		"span_id":      spanCtx.SpanID().String(),
		"trace_id":     spanCtx.TraceID().String(),
		"trace_flags":  spanCtx.TraceFlags().String(),
		"service.name": globalAPMService,
	}
}

// SetGlobalAPMService sets the global APM service name
func SetGlobalAPMService(serviceName string) {
	globalAPMService = serviceName
}
