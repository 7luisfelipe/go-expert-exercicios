package rest

import (
	"encoding/json"
	"fmt"
	"modcleanarch/app/application/service"
	"modcleanarch/app/domain/usecase"
	"modcleanarch/app/infrastructure/database"
	databaseadapter "modcleanarch/app/infrastructure/databaseAdapter"
	"net/http"
)

type RestApi struct {
	PedidoUseCase service.IProdutoService
}

func (controller *RestApi) BuscarPedidos(w http.ResponseWriter, r *http.Request) {
	//Dependências
	controller.PedidoUseCase = &usecase.ProdutoUseCase{
		PedidoRepository: &database.PedidoRepository{
			Conn: &databaseadapter.MariaDbConectar{},
		},
	}

	//Consulta os pedidos
	pedidos, err := controller.PedidoUseCase.ListarPedidos()
	if err != nil {
		fmt.Println("Falha ao consultar pedidos -> REST:")
		fmt.Println(err)
		println()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedidos)
}
