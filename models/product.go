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
	Category           string  `json:"category" gorm:"type:product_category;default:'Food'"`
	// TODO: precisamos ver depois quais serão as categorias, por enquanto este enum é apenas um exemplo

	Market Market `json:"market" gorm:"foreignKey:MarketId;references:ID"`
}

var Products []Product

func (Product) TableName() string {
	return "tbProducts"
}
