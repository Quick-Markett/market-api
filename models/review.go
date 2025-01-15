package models

import "time"

type Review struct {
	ID        uint      `json:"id"`
	ProductId uint      `json:"product_id"`
	OrderId   uint      `json:"order_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`

	Product Product `gorm:"foreignKey:ProductId"`
	Order   Order   `gorm:"foreignKey:OrderId"`
}

var Reviews []Review

func (Review) TableName() string {
	return "tbReviews"
}
