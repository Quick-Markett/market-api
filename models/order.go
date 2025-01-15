package models

import "time"

type Order struct {
	ID              uint      `json:"id"`
	UserId          uint      `json:"user_id"`
	MarketId        uint      `json:"market_id"`
	Status          string    `json:"status"`
	TotalPrice      float64   `json:"total_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeliveryAddress string    `json:"delivery_address"`
	Reviews         []Review  `json:"reviews" gorm:"foreignKey:OrderId"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderId"`
	Market     Market      `json:"market" gorm:"foreignKey:MarketId;references:ID"`
	User       User        `json:"user" gorm:"foreignKey:UserId;references:ID"`
}

var Orders []Order

func (Order) TableName() string {
	return "tbOrders"
}
