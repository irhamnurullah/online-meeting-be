package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Address     []Address `gorm:"foreignKey:UserId"`
	RoomCreated []Room    `gorm:"foreignKey:HostID"`
}
