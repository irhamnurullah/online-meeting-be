package models

import (
	"time"
)

type Education struct {
	ID             uint
	UniversityName string
	Major          string
	EndYear        time.Time
	UserID         uint
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"-"`

	Host User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
