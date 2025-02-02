package models

import "time"

type User struct {
	ID             uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string     `json:"name" gorm:"not null"`
	Email          string     `json:"email" gorm:"not null;unique"`
	Password       string     `json:"password" gorm:"not null"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	City           string     `json:"city"`
	State          string     `json:"state"`
	ProfilePicture string     `json:"profile_picture"`
	CreatedAt      time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"-"`
	IsActive       bool       `json:"is_active" gorm:"default:true"`
	GoogleId       *string    `json:"google_id" gorm:"uniqueIndex"`
	DeletedAt      *time.Time `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "tbUsers"
}
