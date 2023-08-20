package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/models"
)

func (h Handler) LoadUserTemplate(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "users/users.tmpl", gin.H{
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
func (h Handler) FindUsers(ctx *gin.Context) {
	var users []models.User
	h.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new user
func (h Handler) CreateUser(ctx *gin.Context) {
	// Validate request
	var req models.UserCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Name: req.Name, Address: req.Address, Email: req.Email, PhoneNumber: req.PhoneNumber}
	h.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
// Find a user
func (h Handler) FindUser(ctx *gin.Context) { // Get model if exist
	var user models.User

	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update a user
func (h Handler) UpdateUser(ctx *gin.Context) {
	// Get model if exist
	var user models.User
	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate request
	var req models.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Model(&user).Updates(req)

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func (h Handler) DeleteUser(ctx *gin.Context) {
	// Get model if exist
	var user models.User
	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	h.DB.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
