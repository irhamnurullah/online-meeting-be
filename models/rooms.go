package models

import "time"

type Rooms struct {
	ID         uint
	Code       string `gorm:"uniqueIndex"`
	RoomName   string
	ScheduleID uint
	Schedule   ScheduleAppointment `gorm:"foreignKey:ScheduleID" json:"-"`
	CreatedAt  time.Time           `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time           `gorm:"autoUpdateTime" json:"-"`
}
