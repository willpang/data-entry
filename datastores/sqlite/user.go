package sqlite

import (
	"context"
	"errors"

	"github.com/willpang/data-entry/models"
	"gorm.io/gorm"
)

type UserDatastore struct {
	DB *gorm.DB
}

func (d UserDatastore) GetAllUsers(ctx context.Context) []models.User {
	var users []models.User

	d.DB.Find(&users)

	return users
}

func (d UserDatastore) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	if err := d.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &models.User{}, errors.New("Record not found!")
	}

	return &user, nil
}

func (d UserDatastore) CreateUser(ctx context.Context, req models.UserCreateRequest) (*models.User, error) {
	user := models.User{
		Name:        req.Name,
		Address:     req.Address,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}
	d.DB.Create(&user)

	return &user, nil
}

func (d UserDatastore) UpdateUser(ctx context.Context, id string, req models.UserUpdateRequest) (*models.User, error) {
	var user models.User
	if err := d.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &models.User{}, errors.New("Record not found!")
	}

	d.DB.Model(&user).Updates(req)

	return &user, nil
}

func (d UserDatastore) DeleteUser(ctx context.Context, id string) error {
	var user models.User
	if err := d.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return errors.New("Record not found!")
	}

	d.DB.Delete(&user)

	return nil
}
