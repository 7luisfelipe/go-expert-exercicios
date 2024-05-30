package service

import (
	"modcleanarch/app/application/dto"
	"modcleanarch/app/domain/entity"
)

type IProdutoService interface {
	ListarPedidos() ([]entity.Pedido, error)
	CriarPedido(inputDto *dto.CriarPedidoDto) error
}
