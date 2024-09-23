package services

import (
	"errors"
	"fmt"
	"order-api/models"
	"order-api/utils"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type OrderService struct {
	DB *gorm.DB
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := s.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) GetOrderByID(id string) (*models.Order, error) {
	var order models.Order
	if err := s.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *OrderService) GetOrdersByUserID(userID int) ([]models.Order, error) {
	var orders []models.Order
	if err := s.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	exists, err := utils.CheckUserExists(order.UserID)
	if err != nil {
		return errors.New("failed to verify user ID")
	}
	if !exists {
		return errors.New("invalid user ID")
	}

	if err := validate.Struct(order); err != nil {
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

	if err := s.DB.Create(order).Error; err != nil {
		return errors.New("failed to create order")
	}

	return nil
}

func (s *OrderService) UpdateOrder(id string, order *models.Order) (*models.Order, error) {
	var existingOrder models.Order
	if err := s.DB.First(&existingOrder, id).Error; err != nil {
		return nil, errors.New("order not found")
	}

	if order.UserID != 0 {
		existingOrder.UserID = order.UserID
	}
	if order.ItemDescription != "" {
		existingOrder.ItemDescription = order.ItemDescription
	}
	if order.ItemQuantity != 0 {
		existingOrder.ItemQuantity = order.ItemQuantity
	}
	if order.ItemPrice != 0 {
		existingOrder.ItemPrice = order.ItemPrice
	}
	if order.TotalValue != 0 {
		existingOrder.TotalValue = order.TotalValue
	}

	if err := s.DB.Save(&existingOrder).Error; err != nil {
		return nil, errors.New("failed to update order")
	}

	return &existingOrder, nil
}

func (s *OrderService) DeleteOrder(id string) error {
	if err := s.DB.Delete(&models.Order{}, id).Error; err != nil {
		return errors.New("order not found")
	}
	return nil
}
