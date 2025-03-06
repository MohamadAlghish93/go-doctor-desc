package routes

import (
	"doc-desc/controllers"
	"doc-desc/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.POST("/user/register", controllers.RegisterUser)
	r.POST("/user/login", controllers.LoginUser)
	// r.GET("/user/list", controllers.ListUsers)

	// middlewares
	// Protected routes
	protected := r.Group("/")
	//protected.Use(middlewares.AuthMiddleware())
	protected.Use(middlewares.AuthMiddleware("admin", "user"))
	protected.GET("/user/list", controllers.ListUsers)
	protected.POST("/user/register", controllers.RegisterUser)

	protected.POST("/receipt/add", controllers.AddReceipt)
	protected.GET("/receipt/getPdf", controllers.GenerateReceiptPDF)

	return r
}
