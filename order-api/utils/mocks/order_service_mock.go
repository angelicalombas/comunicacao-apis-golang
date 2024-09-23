package mocks

import (
	"order-api/models"

	"github.com/stretchr/testify/mock"
)

type OrderServiceMock struct {
	mock.Mock
}

func (m *OrderServiceMock) GetAllOrders() ([]models.Order, error) {
	args := m.Called()
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *OrderServiceMock) GetOrderByID(id string) (*models.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *OrderServiceMock) GetOrdersByUserID(userID int) ([]models.Order, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *OrderServiceMock) CreateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *OrderServiceMock) UpdateOrder(id string, order *models.Order) (*models.Order, error) {
	args := m.Called(id, order)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *OrderServiceMock) DeleteOrder(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
