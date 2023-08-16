package models

type User struct {
	ID          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
