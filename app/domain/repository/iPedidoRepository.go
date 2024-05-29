package repository

import "modcleanarch/app/domain/entity"

type IPedidoRepository interface {
	BuscarTodosPedidos() (*[]entity.Pedido, error)
}
