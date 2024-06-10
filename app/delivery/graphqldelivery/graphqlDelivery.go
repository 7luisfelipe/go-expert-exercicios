package graphqldelivery

import (
	"errors"
	"fmt"
	"modcleanarch/app/domain/usecase"
	"modcleanarch/app/infrastructure/database"
	databaseadapter "modcleanarch/app/infrastructure/databaseAdapter"

	"github.com/graphql-go/graphql"
)

var PedidoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pedido",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"numeroPedido": &graphql.Field{
				Type: graphql.Int,
			},
			"nomeProduto": &graphql.Field{
				Type: graphql.String,
			},
			"quantidade": &graphql.Field{
				Type: graphql.Int,
			},
			"precoUnitario": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"pedidos": &graphql.Field{
				Type:    graphql.NewList(PedidoType),
				Resolve: BuscarPedidos,
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: RootQuery,
	},
)

type GraphQlServer struct {
	ListarPedidosUseCase usecase.ListarPedidosUseCase
}

func BuscarPedidos(param graphql.ResolveParams) (interface{}, error) {
	//Dependências
	graphQlServer := GraphQlServer{
		ListarPedidosUseCase: usecase.ListarPedidosUseCase{
			PedidoRepository: &database.PedidoRepository{
				Conn: &databaseadapter.MariaDbConectar{},
			},
		},
	}

	//Consulta os pedidos
	pedidos, err := graphQlServer.ListarPedidosUseCase.Execute()
	if err != nil {
		fmt.Println("Falha ao consultar pedidos -> GraphQL:")
		fmt.Println(err)
		println()
		return nil, errors.New("falha ao consultar pedidos com GraphQL")
	}

	return pedidos, nil
}

// GraphQL handler
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		return result
	}
	return result
}
