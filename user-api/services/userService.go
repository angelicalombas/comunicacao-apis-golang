package services

import (
	"errors"
	"fmt"
	"strings"
	"user-api/models"
	"user-api/utils"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate
var cpf string

func init() {
	validate = validator.New()
	validate.RegisterValidation("cpf", func(fl validator.FieldLevel) bool {
		valid, cpfFormat := utils.IsValidCPF(fl.Field().String())
		cpf = cpfFormat
		return valid
	})
}

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) CreateUser(user *models.User) error {
	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, err := range validationErrors {
			errorMessages[err.Field()] = err.Tag()
		}
		var sb strings.Builder
		for field, tag := range errorMessages {
			sb.WriteString(fmt.Sprintf("%s: %s, ", field, tag))
		}
		errorMsg := strings.TrimRight(sb.String(), ", ")
		return errors.New(errorMsg)
	}

	user.CPF = cpf

	var existingUser models.User
	if err := s.DB.Where("cpf = ?", user.CPF).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("error when checking CPF")
		}
	} else {
		return errors.New("CPF already registered")
	}

	if err := s.DB.Create(user).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	var existingUser models.User
	if err := s.DB.First(&existingUser, id).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.PhoneNumber != "" {
		existingUser.PhoneNumber = user.PhoneNumber
	}
	if user.CPF != "" {
		existingUser.CPF = user.CPF
	}

	if err := validate.Struct(existingUser); err != nil {
		return nil, err
	}

	if err := s.DB.Save(&existingUser).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return &existingUser, nil
}

func (s *UserService) DeleteUser(id string) error {
	if err := s.DB.Delete(&models.User{}, id).Error; err != nil {
		return errors.New("user not found")
	}
	return nil
}
