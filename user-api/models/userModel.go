package models

import (
	"time"
)

type User struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" validate:"required"`
	CPF         string     `json:"cpf" validate:"required,cpf"`
	Email       string     `json:"email" validate:"required,email"`
	PhoneNumber string     `json:"phone_number" validate:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
}

type UserRequest struct {
	Name        string `json:"name" validate:"required"`
	CPF         string `json:"cpf" validate:"required,cpf"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
