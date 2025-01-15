package models

type Product struct {
	Id                 int     `json:"id"`
	MarketId           int     `json:"market_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	UnitPrice          float64 `json:"unit_price"`
	Stock              int     `json:"stock"`
	ProductImage       string  `json:"product_image"`
	IsActive           bool    `json:"is_active"`
	
	// TODO: adicioanr campo de categoria futuramente
	// TODO: adicionar relação com mercado futuramente
}

var Products []Product
