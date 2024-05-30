package webhttp

import (
	"encoding/json"
	"log"
	"modcleanarch/app/delivery/graphqldelivery"
	"net/http"
)

func WebGraphQL() {
	http.HandleFunc("POST /graphql/v1/listar", func(w http.ResponseWriter, r *http.Request) {
		var p struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "Requisição inválida", http.StatusBadRequest)
			return
		}

		//result := handler.ExecuteQuery(p.Query, handler.Schema)
		result := graphqldelivery.ExecuteQuery(p.Query, graphqldelivery.Schema)
		json.NewEncoder(w).Encode(result)
	})

	log.Println("Server GhQL rodando na porta: 8083")
	http.ListenAndServe(":8083", nil)
}
