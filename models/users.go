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

	ProfileCreated    []Profile    `gorm:"foreignKey:UserID"`
	ExperienceCreated []Experience `gorm:"foreignKey:UserID"`
	EducationCreated  []Education  `gorm:"foreignKey:UserID"`
}
