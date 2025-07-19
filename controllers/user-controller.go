package controllers

import (
	"net/http"

	"online-meeting/config"

	"github.com/gin-gonic/gin"
)

type Address struct {
	ID      uint   `json:"id"`
	Address string `json:"address"`
}

// ini adalah result yang diharapkan
type UserWithAddresses struct {
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Addresses []Address `json:"addresses"`
}

// kenapa flat karena  SQL tidak “menggabungkan” banyak row ke dalam 1 objek JSON seperti yang biasa kita lihat di frontend.
type flatUserWithAddress struct {
	UserID    uint
	Name      string
	Email     string
	AddressID *uint
	Address   *string
}

func GetUsers(c *gin.Context) {
	var flatResults []flatUserWithAddress
	db := config.DB

	query := `
	SELECT u.id AS user_id, u.name, u.email,
	       a.id AS address_id, a.address
	FROM users u
	LEFT JOIN addresses a ON u.id = a.user_id
	`

	if err := db.Raw(query).Scan(&flatResults).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userMap := make(map[uint]*UserWithAddresses)
	for _, row := range flatResults {
		if _, exists := userMap[row.UserID]; !exists {
			userMap[row.UserID] = &UserWithAddresses{
				UserID:    row.UserID,
				Name:      row.Name,
				Email:     row.Email,
				Addresses: []Address{},
			}
		}
		if row.AddressID != nil && row.Address != nil {
			userMap[row.UserID].Addresses = append(userMap[row.UserID].Addresses, Address{
				ID:      *row.AddressID,
				Address: *row.Address,
			})
		}
	}

	// Convert map to slice
	var users []UserWithAddresses
	for _, u := range userMap {
		users = append(users, *u)
	}

	c.JSON(http.StatusOK, users)
}
