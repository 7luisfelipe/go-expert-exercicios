package dto

type CriarPedidoDto struct {
	NumeroPedido  int     `json:"numeroPedido" validate:"required,min=1"`
	NomeProduto   string  `json:"nomeProduto" validate:"required,max(100)"`
	Quantidade    int     `json:"quantidade" validate:"required,min=1"`
	PrecoUnitario float64 `json:"precoUnitario" validate:"required,min=0.1"`
}
