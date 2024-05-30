package restdelivery

import (
	"encoding/json"
	"fmt"
	"io"
	"modcleanarch/app/application/dto"
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

func (controller *RestApi) CriarPedido(w http.ResponseWriter, r *http.Request) {
	// Dependências
	controller.PedidoUseCase = &usecase.ProdutoUseCase{
		PedidoRepository: &database.PedidoRepository{
			Conn: &databaseadapter.MariaDbConectar{},
		},
	}

	//Dados recebidos
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Faha ao ler dados recebidos - CriarPedido: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var input dto.CriarPedidoDto
	err = json.Unmarshal(body, &input)
	if err != nil {
		fmt.Println("Faha ao passar dados para struct - CriarPedido: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = controller.PedidoUseCase.CriarPedido(&input)

	if err != nil {
		fmt.Println("Faha ao criar produto - CriarProduto: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
