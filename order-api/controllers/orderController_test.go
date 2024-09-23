package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"order-api/models"
	"order-api/utils/mocks"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOrdersSuccess(t *testing.T) {
	mockService := new(mocks.OrderServiceMock)
	mockService.On("GetAllOrders").Return([]models.Order{}, nil)

	router := gin.Default()
	router.GET("/orders", func(c *gin.Context) {
		service := mockService
		orders, err := service.GetAllOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	})

	req, _ := http.NewRequest("GET", "/orders", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetOrderByIDSuccess(t *testing.T) {
	mockService := new(mocks.OrderServiceMock)
	mockService.On("GetOrderByID", "1").Return(&models.Order{}, nil)

	router := gin.Default()
	router.GET("/orders/:id", func(c *gin.Context) {
		service := mockService
		order, err := service.GetOrderByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	})

	req, _ := http.NewRequest("GET", "/orders/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetOrdersByUserIDSuccess(t *testing.T) {
	mockService := new(mocks.OrderServiceMock)
	mockService.On("GetOrdersByUserID", 1).Return([]models.Order{}, nil)

	router := gin.Default()
	router.GET("/users/:id/orders", func(c *gin.Context) {
		service := mockService
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		orders, err := service.GetOrdersByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	})

	req, _ := http.NewRequest("GET", "/users/1/orders", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestCreateOrderSuccess(t *testing.T) {
	mockService := new(mocks.OrderServiceMock)
	order := models.Order{UserID: 1, ItemDescription: "Item", ItemQuantity: 1, ItemPrice: 10.0, TotalValue: 10.0}
	mockService.On("CreateOrder", mock.AnythingOfType("*models.Order")).Return(nil)

	router := gin.Default()
	router.POST("/orders", func(c *gin.Context) {
		service := mockService
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := service.CreateOrder(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, order)
	})

	orderJSON, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(orderJSON))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateOrderSuccess(t *testing.T) {
	mockService := new(mocks.OrderServiceMock)
	order := models.Order{UserID: 1, ItemDescription: "Updated Item", ItemQuantity: 2, ItemPrice: 20.0, TotalValue: 40.0}
	mockService.On("UpdateOrder", "1", &order).Return(&order, nil)

	router := gin.Default()
	router.PUT("/orders/:id", func(c *gin.Context) {
		service := mockService
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedOrder, err := service.UpdateOrder(c.Param("id"), &order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedOrder)
	})

	orderJSON, _ := json.Marshal(order)
	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(orderJSON))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteOrderSuccess(t *testing.T) {
	mockService := new(mocks.OrderServiceMock)
	mockService.On("DeleteOrder", "1").Return(nil)

	router := gin.Default()
	router.DELETE("/orders/:id", func(c *gin.Context) {
		service := mockService
		if err := service.DeleteOrder(c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
	})

	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
