package models

import "time"

type Review struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`

	Product Product `gorm:"foreignKey:ProductId"`
}

var Reviews []Review
