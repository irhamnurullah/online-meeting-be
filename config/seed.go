package config

import (
	"log"
	"online-meeting/models"
)

func SeedData() {
	users := []models.User{
		{
			Name:  "Irham Nurullah",
			Email: "irham@example.com",
			Address: []models.Address{
				{Address: "Jalan Merdeka 1"},
				{Address: "Jalan Sudirman 2"},
			},
		},
		{
			Name:  "Dewi Ayu",
			Email: "dewi@example.com",
			Address: []models.Address{
				{Address: "Jalan Kenangan"},
			},
		},
	}

	for _, user := range users {
		if err := DB.Create(&user).Error; err != nil {
			log.Println("Gagal insert user:", err)
		}
	}

}
