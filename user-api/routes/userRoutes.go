package routes

import (
	"user-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", controllers.GetUsers(db))
	r.GET("/users/:id", controllers.GetUserByID(db))
	r.POST("/users", controllers.CreateUser(db))
	r.PUT("/users/:id", controllers.UpdateUser(db))
	r.DELETE("/users/:id", controllers.DeleteUser(db))
}
