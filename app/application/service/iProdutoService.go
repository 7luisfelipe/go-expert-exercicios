package service

import "modcleanarch/app/domain/entity"

type IProdutoService interface {
	ListarPedidos() ([]entity.Pedido, error)
}
