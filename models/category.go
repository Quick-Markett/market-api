package models

import "time"

type Category struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"not null"`
	MarketId    int        `json:"market_id"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"-"`

	Market   Market    `json:"market" gorm:"foreignKey:MarketId;references:ID"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryId;references:ID"`
}

var Categories []Category

func (Category) TableName() string {
	return "tbCategories"
}
