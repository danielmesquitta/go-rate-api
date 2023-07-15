package router

import (
	"github.com/danielmesquitta/go-rate-api/docs"
	"github.com/danielmesquitta/go-rate-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Auth
		v1.POST("/login", handler.LoginHandler)
	}
	{
		// Users
		v1.POST("/user", handler.CreateUserHandler)
		v1.PUT("/user", handler.UpdateUserHandler)
		v1.DELETE("/user", handler.DeleteUserHandler)
		v1.GET("/user", handler.ShowUserHandler)
		v1.GET("/users", handler.ListUsersHandler)
		v1.GET("/me", handler.RequireAuth, handler.ShowMeHandler)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
