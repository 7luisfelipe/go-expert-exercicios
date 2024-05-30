package webhttp

import (
	"log"
	"modcleanarch/app/delivery/restdelivery"
	"net/http"
)

func WebRest() {
	apiRest := restdelivery.RestApi{}

	//localhost:8081/rest/v1/listar
	http.HandleFunc("GET /rest/v1/listar", apiRest.BuscarPedidos)
	http.HandleFunc("POST /rest/v1/cadastrarPedido", apiRest.CriarPedido)

	log.Printf("Server REST rodando na porta: 8081")
	http.ListenAndServe(":8081", nil)
}
