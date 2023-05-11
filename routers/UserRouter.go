package routers

import (
	session "documentation/middlewares"
	"documentation/services"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(rg *gin.RouterGroup) {
	usersRoute := rg.Group("/users", session.SetSession())

	usersRoute.GET("/", services.FindAllUsers)
	usersRoute.GET("/:email", services.FindByUserEmail)
	usersRoute.POST("/", services.CreateUser)
	usersRoute.POST("/login", services.LoginUser)
	usersRoute.GET("/check", services.CheckUserSession)

	usersRoute.Use(session.AuthSession())
	{
		usersRoute.DELETE("/:id", services.DeleteUser)
		usersRoute.PUT("/:id", services.UpdateUser)
		usersRoute.GET("/logout", services.LogoutUser)
	}
}
