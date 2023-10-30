package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/models"
)

func (h Handler) LoadUserTemplate(ctx *gin.Context) {
	users := h.UserDatastore.GetAllUsers(ctx)

	ctx.HTML(http.StatusOK, "users/users.tmpl", gin.H{
		"Title":    "Pelanggan",
		"UserData": users,
	})
}

// GET /users
// Get all users
func (h Handler) FindUsers(ctx *gin.Context) {
	users := h.UserDatastore.GetAllUsers(ctx)

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
	user, err := h.UserDatastore.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": *user})
}

// GET /users/:id
// Find a user
func (h Handler) FindUser(ctx *gin.Context) { // Get model if exist
	user, err := h.UserDatastore.GetUser(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": *user})
}

// PATCH /users/:id
// Update a user
func (h Handler) UpdateUser(ctx *gin.Context) {
	// Validate request
	var req models.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user
	user, err := h.UserDatastore.UpdateUser(ctx, ctx.Param("id"), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": *user})
}

// DELETE /users/:id
// Delete a user
func (h Handler) DeleteUser(ctx *gin.Context) {
	if err := h.UserDatastore.DeleteUser(ctx, ctx.Param("id")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
