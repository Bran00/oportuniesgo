package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Bran00/oportuniesgo/handler"
)

import swaggerfiles "github.com/swaggo/files"
import ginSwagger "github.com/swaggo/gin-swagger"
import docs "github.com/Bran00/oportuniesgo/docs"


func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ShowListOpeningHandler)
	}
	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
