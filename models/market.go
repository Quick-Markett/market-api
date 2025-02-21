package models

import "time"

type Market struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	Slug            string     `json:"slug" gorm:"not null;unique"`
	Name            string     `json:"name" gorm:"not null"`
	Email           string     `json:"email" gorm:"not null;unique"`
	PhoneNumber     string     `json:"phone_number"`
	Address         string     `json:"address"`
	City            string     `json:"city"`
	State           string     `json:"state"`
	ZipCode         string     `json:"zip_code"`
	Description     string     `json:"description"`
	LogoUrl         string     `json:"logo_url"`
	DeliveryPrice   float64    `json:"delivery_price" gorm:"default:0"`
	DeliveryMinTime float64    `json:"delivery_min_time" gorm:"default:0"`
	DeliveryMaxTime float64    `json:"delivery_max_time" gorm:"default:30"`
	CreatedAt       time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       *time.Time `json:"updated_at" gorm:"-"`
	IsActive        bool       `json:"is_active" gorm:"default:true"`
	DeletedAt       *time.Time `json:"deleted_at" gorm:"index"`
}

func (Market) TableName() string {
	return "tbMarkets"
}
