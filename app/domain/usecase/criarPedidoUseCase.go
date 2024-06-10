package usecase

import (
	"modcleanarch/app/application/dto"
	"modcleanarch/app/domain/entity"
	"modcleanarch/app/domain/repository"
)

type CriarPedidoUseCase struct {
	PedidoRepository repository.IPedidoRepository
}

func (useCase *CriarPedidoUseCase) Execute(inputDto *dto.CriarPedidoDto) error {

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
