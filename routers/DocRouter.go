package routers

import (
	session "documentation/middlewares"
	"documentation/services"

	"github.com/gin-gonic/gin"
)

func AddDocRouter(rg *gin.RouterGroup) {
	docsRoute := rg.Group("/docs", session.SetSession())

	// docsRoute.GET("/", services.FindAllUsers)

	docsRoute.Use(session.AuthSession())
	{
		docsRoute.POST("/createDoc", services.CreateDoc)
		// 	docsRoute.DELETE("/:id", services.DeleteUser)
		// 	docsRoute.PUT("/:id", services.UpdateUser)
	}
}
