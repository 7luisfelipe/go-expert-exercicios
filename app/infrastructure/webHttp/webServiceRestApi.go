package webhttp

import (
	"modcleanarch/app/delivery/rest"
	"net/http"
)

func Routes() {
	apiRest := rest.RestApi{}

	//localhost:8081/rest/v1/listar
	http.HandleFunc("GET /rest/v1/listar", apiRest.BuscarPedidos)

	http.ListenAndServe(":8081", nil)
}
