package models

import "time"

type Product struct {
	ID                 uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	MarketId           int        `json:"market_id"`
	ProductName        string     `json:"product_name"`
	ProductDescription string     `json:"product_description"`
	UnitPrice          float64    `json:"unit_price"`
	Stock              int        `json:"stock"`
	ProductImage       string     `json:"product_image"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at" gorm:"-"`
	IsActive           bool       `json:"is_active"`
	// Category           string  `json:"category" gorm:"type:product_category;default:'Food'"`
	// TODO: precisamos ver depois quais serão as categorias, por enquanto este enum é apenas um exemplo

	Market Market `json:"market" gorm:"foreignKey:MarketId;references:ID"`
}

var Products []Product

func (Product) TableName() string {
	return "tbProducts"
}
