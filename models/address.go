package models

import "time"

type Address struct {
	ID        uint `gorm:"primaryKey"`
	UserId    string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
}
