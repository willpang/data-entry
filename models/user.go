package models

type User struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type UserUpdateRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
