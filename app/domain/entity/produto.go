package entity

type Produto struct {
	Base
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}
