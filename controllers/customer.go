package controllers

import (
	"go-auth/models"

	"go-auth/utils"

	"github.com/gin-gonic/gin"
)

func GetDetailCustomer(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var req models.Customer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer

	models.DB.Where("id = ?", req.ID).First(&customer)

	c.JSON(200, gin.H{"success": "detail Customer", "data": customer})
}

func InsertCustomer(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var req models.Customer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingCustomer models.Customer

	models.DB.Where("PhoneNumber = ?", req.PhoneNumber).First(&existingCustomer)

	if existingCustomer.ID != 0 {
		c.JSON(400, gin.H{"error": "customer already exists"})
		return
	}

	models.DB.Create(&req)

	c.JSON(200, gin.H{"success": "Customer created"})
}

func UpdateCustomer(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var req models.Customer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.ID == 0 {
		c.JSON(401, gin.H{"error": "Customer id not found"})
		return
	}

	models.DB.Model(&req).Updates(models.Customer{PhoneNumber: req.PhoneNumber, Name: req.Name})

	c.JSON(200, gin.H{"success": "Customer updated"})
}

func DeleteCustomer(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var req models.Customer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&req)

	c.JSON(200, gin.H{"success": "Customer deleted"})
}

func SearchCustomer(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var req models.Customer

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	models.DB.Where(&models.Customer{PhoneNumber: req.PhoneNumber, Name: req.Name})

	c.JSON(200, gin.H{"success": "Customer deleted"})
}
