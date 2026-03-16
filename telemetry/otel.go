package telemetry

import (
	"context"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func Init(ctx context.Context) (func(context.Context) error, error) {
	serviceName := getenv("OTEL_SERVICE_NAME", "sample-api")
	otlpEndpoint := getenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://172.15.1.100:4318")

	res, err := resource.New(
		ctx,
		resource.WithFromEnv(),
		resource.WithAttributes(attribute.String("service.name", serviceName)),
	)
	if err != nil {
		return nil, err
	}

	traceCtx, traceCancel := context.WithTimeout(ctx, 10*time.Second)
	defer traceCancel()

	traceExporter, err := otlptracehttp.New(
		traceCtx,
		otlptracehttp.WithEndpointURL(otlpEndpoint),
	)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	metricCtx, metricCancel := context.WithTimeout(ctx, 10*time.Second)
	defer metricCancel()

	metricExporter, err := otlpmetrichttp.New(
		metricCtx,
		otlpmetrichttp.WithEndpointURL(otlpEndpoint),
	)
	if err != nil {
		return nil, err
	}

	mp := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
	)
	otel.SetMeterProvider(mp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return func(shutdownCtx context.Context) error {
		var shutdownErr error
		if err := mp.Shutdown(shutdownCtx); err != nil {
			shutdownErr = err
		}
		if err := tp.Shutdown(shutdownCtx); err != nil && shutdownErr == nil {
			shutdownErr = err
		}
		return shutdownErr
	}, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
