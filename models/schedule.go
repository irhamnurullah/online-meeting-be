package models

import "time"

type ScheduleAppointment struct {
	ID            uint
	StartDate     time.Time
	EndDate       time.Time
	AppointmentID uint
	Appointment   Appointment `gorm:"foreignKey:AppointmentID" json:"-"`
	CreatedAt     time.Time   `gorm:"autoCreateTime" json:"-"`
	UpdatedAt     time.Time   `gorm:"autoUpdateTime" json:"-"`
}
