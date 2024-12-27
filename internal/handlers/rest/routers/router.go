package routers

import (
	"github.com/baking-bread/boutique/internal/handlers/rest/controllers"
	"github.com/baking-bread/boutique/internal/handlers/rest/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.DatabaseMiddleware())

	public := router.Group("/api")
	{
		public.GET("/ping", controllers.GetPing)
	}
	return router
}
