package models

type Product struct {
	ID                 uint    `json:"id"`
	MarketId           uint    `json:"market_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	UnitPrice          float64 `json:"unit_price"`
	Stock              int     `json:"stock"`
	ProductImage       string  `json:"product_image"`
	IsActive           bool    `json:"is_active"`

	Market Market `json:"market" gorm:"foreignKey:MarketId;references:ID"`

	// TODO: adicioanr campo de categoria futuramente
}

var Products []Product

func (Product) TableName() string {
	return "tbProducts"
}
