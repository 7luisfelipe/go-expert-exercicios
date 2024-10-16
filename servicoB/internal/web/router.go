package web

import (
	"log"
	"modapilab1/internal/controller"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Router struct {
	FindDataHandler controller.FindDataController
}

func (r *Router) StartRouter() {
	/*
		http.HandleFunc("/", r.FindDataHandler.FindData)
		http.ListenAndServe(":8081", nil)
	*/
	// Configurar o cliente HTTP com middleware OpenTelemetry
	// Configurar o manipulador HTTP com middleware OpenTelemetry
	handler := otelhttp.NewHandler(http.HandlerFunc(r.FindDataHandler.FindData), "FindData")

	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
