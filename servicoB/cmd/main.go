package main

import (
	"context"
	"log"
	"modapilab1/internal/web"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func main() {
	/*
		r := web.Router{}
		r.StartRouter()
	*/
	// URL do servidor Zipkin

	// URL do servidor Zipkin
	zipkinURL := "http://zipkin:9411/api/v2/spans"

	// Criar o exportador Zipkin
	exporter, err := zipkin.New(zipkinURL)
	if err != nil {
		log.Fatalf("failed to create zipkin exporter: %v", err)
	} else {
		log.Println("Zipkin exporter configurado com sucesso")
	}

	/*
		// Criar o tracer provider
		tp := sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
		)
	*/
	// Configurar o provedor de trace
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("servico-b"),
		)),
	)
	defer func() { _ = tp.Shutdown(context.Background()) }()

	// Configurar o tracer global
	otel.SetTracerProvider(tp)

	// Configurar o propagador global
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	// Iniciar o roteador
	r := web.Router{}
	r.StartRouter()
}
