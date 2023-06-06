package repositories

import (
	"errors"
	"fmt"
	"goselflearn/internal/initializers"
	"goselflearn/internal/models"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct{}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	dbResult := initializers.DB.Model(&models.User{}).Where(models.User{Email: email}).Preload("Posts").First(&user)
	if dbResult.Error != nil {
		fmt.Println("find by email error: ", dbResult.Error.Error())
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, dbResult.Error
	}
	return &user, nil
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	dbResult := initializers.DB.Model(&models.User{}).Create(user)

	if errors.Is(dbResult.Error, gorm.ErrDuplicatedKey) {
		if emailExists := strings.Contains(strings.ToLower(dbResult.Error.Error()), "email"); emailExists {
			return fmt.Errorf("EMAIL_EXIST")
		}
		return dbResult.Error
	}

	return nil
}

func (ur *UserRepository) FindUserById(id uint64) (*models.User, error) {
	var user models.User
	dbResult := initializers.DB.Model(&models.User{}).Where("id = ?", id).Preload("Posts").First(&user)
	if dbResult.Error != nil {
		fmt.Println("find by email error: ", dbResult.Error.Error())
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, dbResult.Error
	}
	return &user, nil
}