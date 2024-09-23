package routes

import (
	"order-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func OrderRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/orders", controllers.GetOrders(db))
	r.GET("/orders/:id", controllers.GetOrderByID(db))
	r.GET("/users/:id/orders", controllers.GetOrdersByUserID(db))
	r.POST("/orders", controllers.CreateOrder(db))
	r.PUT("/orders/:id", controllers.UpdateOrder(db))
	r.DELETE("/orders/:id", controllers.DeleteOrder(db))
}
