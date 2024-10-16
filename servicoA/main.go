package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"servicoa/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func zipCodeHandler(w http.ResponseWriter, r *http.Request) {

	// Parâmetro
	zipCodeParam := r.URL.Query().Get("cep")
	if zipCodeParam == "" || len(zipCodeParam) != 8 {
		w.WriteHeader(422)
		w.Write([]byte(config.INVALID_ZIP_CODE))
		return
	}

	//resp, err := http.Get("http://localhost:8081/?cep=" + zipCodeParam)
	//resp, err := http.Get("http://api-servico-b:8081/?cep=" + zipCodeParam)
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Get("http://api-servico-b:8081/?cep=" + zipCodeParam)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err.Error())
		w.WriteHeader(404)
		w.Write([]byte(config.ZIP_CODE_NOT_FOUND))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(config.ZIP_CODE_NOT_FOUND))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

}

func main() {
	/*
		http.HandleFunc("/", zipCodeHandler)
		http.ListenAndServe(":8080", nil)
	*/
	// Configurar o exportador Prometheus
	//exporter, err := prometheus.New(prometheus.Config{})
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("failed to initialize prometheus exporter: %v", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(exporter),
	)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ServicoA"),
		)),
	)

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(meterProvider)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(zipCodeHandler), "zipCodeHandler"))
	log.Println("Serviço A rodando na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
