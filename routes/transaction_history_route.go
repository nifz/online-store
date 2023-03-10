package routes

import (
	"online-store/controllers"
	"online-store/middlewares"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(c *gin.Engine, controller controllers.TransactionController) {
	c.POST("/checkout", middlewares.AuthMiddleware, controller.CreateTransaction)
	c.GET("/transactions", middlewares.AuthMiddleware, controller.GetUserTransactions)
	c.GET("/transactions/user-transactions", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.GetAllTransactions)
}
