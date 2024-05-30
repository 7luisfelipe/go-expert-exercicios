package usecase

import (
	"modcleanarch/app/application/dto"
	"modcleanarch/app/domain/entity"
	"modcleanarch/app/domain/repository"
)

// Implements IProdutoService
type ProdutoUseCase struct {
	PedidoRepository repository.IPedidoRepository
}

func (useCase *ProdutoUseCase) ListarPedidos() ([]entity.Pedido, error) {
	pedidos, err := useCase.PedidoRepository.BuscarTodosPedidos()
	if err != nil {
		return nil, err
	}
	return *pedidos, nil
}

func (useCase *ProdutoUseCase) CriarPedido(inputDto *dto.CriarPedidoDto) error {

	pedido := entity.Pedido{
		NumeroPedido:  inputDto.NumeroPedido,
		NomeProduto:   inputDto.NomeProduto,
		Quantidade:    inputDto.Quantidade,
		PrecoUnitario: inputDto.PrecoUnitario,
	}

	err := useCase.PedidoRepository.CriarPedido(&pedido)
	if err != nil {
		return err
	}

	return nil
}
