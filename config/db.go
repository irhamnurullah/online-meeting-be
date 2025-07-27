package config

import (
	"fmt"
	"log"
	"online-meeting/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	LoadEnv()

	// dsn := "host=switchyard.proxy.rlwy.net user=postgres password=CYcBzUudHWNrLMIPxAoYjSUHElSJmhdr dbname=railway port=41775 sslmode=disable"

	timezone := GetEnv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_DATABASE"),
		GetEnv("DB_PORT"),
		GetEnv("DB_SSL"),
	)

	if timezone != "" {
		dsn += fmt.Sprintf(" TimeZone=%s", timezone)
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	var autoMigrate = GetEnv("AUTO_MIGRATE")

	// migrasi bila diperlukan
	if autoMigrate == "true" {
		err = database.AutoMigrate(
			&models.User{},
			&models.Rooms{},
			&models.Profile{},
			&models.Education{},
			&models.Experience{},
			&models.Appointment{},
			&models.ScheduleAppointment{},
		)
	}

	if err != nil {
		log.Fatal("failed to migrate: ", err)
	}

	DB = database
}
