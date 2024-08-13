package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	PhoneNumber string `gorm:"unique" json:"phone_number"`
	Name        string `json:"name"`
}
