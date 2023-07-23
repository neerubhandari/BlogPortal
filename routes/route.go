package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/BlogPortal/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/register", controller.Register)

	return router
}