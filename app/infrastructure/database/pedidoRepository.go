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

// Cadastra um pedido
func (repository *PedidoRepository) CriarPedido(produto *entity.Pedido) error {
	db := repository.Conn.GetConn()
	defer db.Close()

	command := "INSERT INTO pedidos (numero_pedido, nome_produto, quantidade, preco_unitario) VALUES(?, ?, ?, ?);"
	_, err := db.Exec(command, produto.NumeroPedido, produto.NomeProduto, produto.Quantidade, produto.PrecoUnitario)
	if err != nil {
		return err
	}

	return nil
}

// Consulta todos os pedidos
func (repository *PedidoRepository) BuscarTodosPedidos() (*[]entity.Pedido, error) {
	db := repository.Conn.GetConn()
	defer db.Close()
	rows, err := db.Query(`
						 SELECT p.numero_pedido,
								p.nome_produto,
								p.quantidade,
								p.preco_unitario
						FROM pedidos p
						ORDER BY p.id ASC;
	`)
	if err != nil {
		return nil, err
	}

	var pedidos []entity.Pedido

	for rows.Next() {

		var p entity.Pedido

		err = rows.Scan(
			&p.NumeroPedido,
			&p.NomeProduto,
			&p.Quantidade,
			&p.PrecoUnitario,
		)
		if err != nil {
			return nil, err
		}

		pedidos = append(pedidos, p)
	}

	return &pedidos, nil
}
