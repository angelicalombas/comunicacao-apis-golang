package controllers

import (
	"net/http"
	"user-api/models"
	"user-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	service := services.UserService{DB: db}
	return func(c *gin.Context) {
		users, err := service.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get a specific user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [get]
func GetUserByID(db *gorm.DB) gin.HandlerFunc {
	service := services.UserService{DB: db}
	return func(c *gin.Context) {
		user, err := service.GetUserByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param UserRequest body models.UserRequest true "UserRequest"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	service := services.UserService{DB: db}
	return func(c *gin.Context) {
		var userRequest models.UserRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.User{
			Name:        userRequest.Name,
			CPF:         userRequest.CPF,
			Email:       userRequest.Email,
			PhoneNumber: userRequest.PhoneNumber,
		}

		if err := service.CreateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update an existing user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param UserRequest body models.UserRequest true "UserRequest"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /users/{id} [put]
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	service := services.UserService{DB: db}
	return func(c *gin.Context) {
		var userRequest models.UserRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.User{
			Name:        userRequest.Name,
			CPF:         userRequest.CPF,
			Email:       userRequest.Email,
			PhoneNumber: userRequest.PhoneNumber,
		}

		updatedUser, err := service.UpdateUser(c.Param("id"), &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedUser)
	}
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	service := services.UserService{DB: db}
	return func(c *gin.Context) {
		if err := service.DeleteUser(c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
