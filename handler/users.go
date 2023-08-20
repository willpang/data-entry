package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/datastore/sqlite"
	"github.com/willpang/data-entry/models"
)

func LoadUserTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "users/users.tmpl", gin.H{
		"Title": "Users Page",
		"UserData": []models.User{
			{
				Name:        "Test 1",
				Address:     "Street 1",
				Email:       "Test1@gmail.com",
				PhoneNumber: "0811234567890",
			},
			{
				Name:        "Test 2",
				Address:     "Street 2",
				Email:       "Test2@gmail.com",
				PhoneNumber: "0821234567890",
			},
		},
	})
}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	sqlite.DB.Find(&users)

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
	sqlite.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	if err := sqlite.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := sqlite.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlite.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := sqlite.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	sqlite.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
