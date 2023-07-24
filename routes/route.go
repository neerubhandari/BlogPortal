package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/BlogPortal/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/register", controller.Register)
	router.POST("/api/login", controller.Login)
	// router.Use(middleware.IsAuthenticate)
	router.POST("/api/post", controller.CreatePost)
	router.GET("/api/post", controller.AllPost)
	router.GET("/api/post/:id", controller.DetailPost)
	router.PUT("/api/post/:id", controller.UpdatePost)
	router.GET("/api/uniquepost", controller.UniquePost)

	return router
}
