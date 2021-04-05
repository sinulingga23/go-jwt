package model

type Product struct {
	ProductId	int 	`json:"productId"`
	CategoryId	string  `json:"categoryProductId"`
	Name		string  `json:"product"`
	Unit		string  `json:"unit"`
	Price		float64 `json:"price"`
	Stock           int     `json:"stock"`
	Audit           Audit   `json:"audit"`
}
