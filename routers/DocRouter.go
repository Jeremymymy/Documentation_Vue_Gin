package routers

import (
	session "documentation/middlewares"
	"documentation/services"

	"github.com/gin-gonic/gin"
)

func AddDocRouter(rg *gin.RouterGroup) {
	docsRoute := rg.Group("/docs", session.SetSession())

	docsRoute.GET("/", services.FindAllUsers)
	// docsRoute.GET("/:email", services.FindByUserEmail)
	// docsRoute.POST("/", services.CreateUser)
	// docsRoute.POST("/login", services.LoginUser)
	// docsRoute.GET("/check", services.CheckUserSession)

	// docsRoute.Use(session.AuthSession())
	// {
	// 	docsRoute.DELETE("/:id", services.DeleteUser)
	// 	docsRoute.PUT("/:id", services.UpdateUser)
	// 	docsRoute.GET("/logout", services.LogoutUser)
	// }
}
