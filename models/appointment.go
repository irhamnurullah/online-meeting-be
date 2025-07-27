package models

import (
	"time"
)

type AppointmentStatus string

const (
	StatusBooking   AppointmentStatus = "booking"   // Saat mentee baru membuat permintaan
	StatusScheduled AppointmentStatus = "scheduled" // Setelah disetujui & dijadwalkan
	StatusDone      AppointmentStatus = "done"      // Setelah meeting selesai
)

type Appointment struct {
	ID        uint
	Objective string
	Metric    string
	Chellenge string
	Status    AppointmentStatus `gorm:"type:varchar(20)" json:"status"`
	IDMentor  uint
	IDMentee  uint
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`

	// Relasi ke Profile
	Mentor Profile `gorm:"foreignKey:IDMentor;references:ID" json:"mentor"`
	Mentee Profile `gorm:"foreignKey:IDMentee;references:ID" json:"mentee"`
}
