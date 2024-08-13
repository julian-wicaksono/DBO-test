package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Detail     string `json:"name"`
	CustomerID uint   `json:"customer_id"`
	Status     int    `json:"status"`
}
