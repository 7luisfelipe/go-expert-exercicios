package entity

type Pedido struct {
	Base
	NumeroPedido int          `json:"numeroPedido"`
	ItensPedido  []ItemPedido `json:"itensPedido"`
}
