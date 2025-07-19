package models

import "time"

type Room struct {
	ID        uint
	Code      string `gorm:"uniqueIndex"`
	RoomName  string
	HostID    uint
	StartTime time.Time
	EndTime   time.Time

	Host User `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
}
