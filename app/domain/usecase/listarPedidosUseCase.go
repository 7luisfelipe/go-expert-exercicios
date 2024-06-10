package usecase

import (
	"modcleanarch/app/domain/entity"
	"modcleanarch/app/domain/repository"
)

// Implements IProdutoService
type ListarPedidosUseCase struct {
	PedidoRepository repository.IPedidoRepository
}

func (useCase *ListarPedidosUseCase) Execute() ([]entity.Pedido, error) {
	pedidos, err := useCase.PedidoRepository.BuscarTodosPedidos()
	if err != nil {
		return nil, err
	}
	return *pedidos, nil
}
