package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/models"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Name: input.Name, Address: input.Address, Email: input.Email, PhoneNumber: input.PhoneNumber}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
