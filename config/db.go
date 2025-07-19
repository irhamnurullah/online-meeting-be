package config

import (
	"log"
	"online-meeting/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)

var DB *gorm.DB

func ConnectionDatabase() {
	config := LoadConfig()

	dsn := config.Database.GetDSN()
	// dsn := "host=localhost user=irhamnurullah password=123 dbname=belajar-go-pg port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	// migrasi bila diperlukan
	if os.Getenv("AUTO_MIGRATE") == "true" {
		err = database.AutoMigrate(
			&models.User{},
			&models.Address{},
			&models.Room{},
		)
	}

	if err != nil {
		log.Fatal("failed to migrate: ", err)
	}

	DB = database
}
