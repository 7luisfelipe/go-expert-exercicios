package repository

import (
	"modcleanarch/app/domain/entity"
)

type IPedidoRepository interface {
	CriarPedido(produto *entity.Pedido) error
	BuscarTodosPedidos() (*[]entity.Pedido, error)
}
