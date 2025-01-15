package models

import "time"

type Order struct {
	Id              int       `json:"id"`
	UserId          int       `json:"user_id"`
	MarketId        int       `json:"market_id"`
	Status          string    `json:"status"`
	TotalPrice      float64   `json:"total_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeliveryAddress string    `json:"delivery_address"`
	Reviews         []Review  `json:"reviews"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
	// TODO: adicionar relação com mercado e usuario aqui
}

var Orders []Order

type OrderItem struct {
	Id         int     `json:"id"`
	OrderId    int     `json:"order_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`

	Product Product `gorm:"foreignKey:ProductId"`
}

var OrderItems []OrderItem
