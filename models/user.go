package models

import "time"

type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	Name         string     `json:"name" gorm:"not null"`
	Email        string     `json:"email" gorm:"not null;unique"`
	PasswordHash string     `json:"password_hash" gorm:"not null"`
	PhoneNumber  string     `json:"phone_number"`
	Address      string     `json:"address"`
	City         string     `json:"city"`
	State        string     `json:"state"`
	ZipCode      string     `json:"zip_code"`
	CreatedAt    time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"-"`
	IsActive     bool       `json:"is_active" gorm:"default:true"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "tbUsers"
}
