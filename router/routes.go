package router

import (
	"github.com/danielmesquitta/go-rate-api/controller"
	"github.com/danielmesquitta/go-rate-api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Users
		v1.POST("/user", controller.CreateUserController)
		v1.PUT("/user", controller.UpdateUserController)
		v1.DELETE("/user", controller.DeleteUserController)
		v1.GET("/user", controller.ShowUserController)
		v1.GET("/users", controller.ListUsersController)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
