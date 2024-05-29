package entity

type ItemPedido struct {
	Quantidade int     `json:"Quantidade"`
	Preco      float64 `json:"preco"`
	Produto    Produto `json:"produto"`
}
