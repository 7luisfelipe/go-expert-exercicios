package database

import (
	"modcleanarch/app/adapters"
	"modcleanarch/app/domain/entity"
)

// Implements IPedidoRepository
type PedidoRepository struct {
	//DB *sql.DB
	Conn adapters.IConectarBanco
}

func (repository *PedidoRepository) BuscarTodosPedidos() (*[]entity.Pedido, error) {
	//(ip.quantidade * prod.preco) AS valor_total_item
	db := repository.Conn.GetConn()
	defer db.Close()
	rows, err := db.Query(`
						 SELECT p.numero_pedido,
								ip.quantidade,
								ip.preco,
								prod.nome AS produto_nome,
								prod.preco
						FROM pedidos p
						JOIN itens_pedido ip ON p.id = ip.pedido_id
						JOIN produtos prod ON ip.produto_id = prod.id
						ORDER BY p.numero_pedido, ip.id;
	`)
	if err != nil {
		return nil, err
	}

	var pedidos []entity.Pedido
	var pedido entity.Pedido
	pedido.NumeroPedido = -1

	for rows.Next() {

		var pd entity.Pedido
		var it entity.ItemPedido

		err = rows.Scan(
			&pd.NumeroPedido,
			&it.Quantidade,
			&it.Preco,
			&it.Produto.Nome,
			&it.Produto.Preco,
		)
		if err != nil {
			return nil, err
		}

		//Valida se é um novo pedido
		if pedido.NumeroPedido != pd.NumeroPedido {

			//Se o NumeroPedido for -1 então é o valor inicial, não um pedido para add na lista
			if pedido.NumeroPedido != -1 {
				//É um pedido válido e já leu todos os itens, então add na lista
				pedidos = append(pedidos, pedido)
			}
			//Novo pedido
			pedido.NumeroPedido = pd.NumeroPedido
		}

		//itens do pedido
		pedido.ItensPedido = append(pedido.ItensPedido, it)
	}

	return &pedidos, nil
}
