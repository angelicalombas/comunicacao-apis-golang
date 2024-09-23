package controllers

import (
	"net/http"
	"order-api/models"
	"order-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Router /orders [get]
func GetOrders(db *gorm.DB) gin.HandlerFunc {
	service := services.OrderService{DB: db}
	return func(c *gin.Context) {
		orders, err := service.GetAllOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Get a specific order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 404 {object} models.ErrorResponse
// @Router /orders/{id} [get]
func GetOrderByID(db *gorm.DB) gin.HandlerFunc {
	service := services.OrderService{DB: db}
	return func(c *gin.Context) {
		order, err := service.GetOrderByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Order not found"})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

// GetOrdersByUserID godoc
// @Summary Get orders by user ID
// @Description Get orders for a specific user by user ID
// @Tags orders
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.Order
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id}/orders [get]
func GetOrdersByUserID(db *gorm.DB) gin.HandlerFunc {
	service := services.OrderService{DB: db}
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
			return
		}
		orders, err := service.GetOrdersByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch orders"})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

// CreateOrder godoc
// @Summary Create a new OrderRequest
// @Description Create a new OrderRequest
// @Tags orders
// @Accept json
// @Produce json
// @Param OrderRequest body models.OrderRequest true "OrderRequest"
// @Success 201 {object} models.Order
// @Failure 400 {object} models.ErrorResponse
// @Router /orders [post]
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	service := services.OrderService{DB: db}
	return func(c *gin.Context) {
		var orderRequest models.OrderRequest
		if err := c.ShouldBindJSON(&orderRequest); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		order := models.Order{
			UserID:          orderRequest.UserID,
			ItemDescription: orderRequest.ItemDescription,
			ItemQuantity:    orderRequest.ItemQuantity,
			ItemPrice:       orderRequest.ItemPrice,
			TotalValue:      orderRequest.TotalValue,
		}

		if err := service.CreateOrder(&order); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, order)
	}
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an existing order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param OrderRequest body models.OrderRequest true "OrderRequest"
// @Success 200 {object} models.Order
// @Failure 400 {object} models.ErrorResponse
// @Router /orders/{id} [put]
func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	service := services.OrderService{DB: db}
	return func(c *gin.Context) {
		var orderRequest models.OrderRequest
		if err := c.ShouldBindJSON(&orderRequest); err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		order := models.Order{
			UserID:          orderRequest.UserID,
			ItemDescription: orderRequest.ItemDescription,
			ItemQuantity:    orderRequest.ItemQuantity,
			ItemPrice:       orderRequest.ItemPrice,
			TotalValue:      orderRequest.TotalValue,
		}

		updatedOrder, err := service.UpdateOrder(c.Param("id"), &order)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedOrder)
	}
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by ID
// @Tags orders
// @Param id path int true "Order ID"
// @Success 200 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /orders/{id} [delete]
func DeleteOrder(db *gorm.DB) gin.HandlerFunc {
	service := services.OrderService{DB: db}
	return func(c *gin.Context) {
		if err := service.DeleteOrder(c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Order not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
	}
}
