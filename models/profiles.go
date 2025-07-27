package models

import (
	"time"
)

type Profile struct {
	ID                 uint `gorm:"primaryKey"`
	Fullname           string
	DateBirth          time.Time
	ProfileDescription string
	ContactPerson      string
	ProfilePicture     string
	IsMentor           bool
	UserID             uint      `gorm:"unique"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"-"`

	Host        User         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Educations  []Education  `gorm:"foreignKey:UserID;references:UserID"`
	Experiences []Experience `gorm:"foreignKey:UserID;references:UserID"`
}
