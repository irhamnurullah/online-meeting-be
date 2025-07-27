package models

import (
	"time"
)

type Experience struct {
	ID              uint
	CurrentPosition string
	CompanyName     string
	StartYear       time.Time
	Skills          string
	Achievement     string
	UserID          uint
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"-"`

	Host User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
