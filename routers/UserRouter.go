package routers

import (
	"documentation/services"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(rg *gin.RouterGroup) {
	usersRoute := rg.Group("/users")

	usersRoute.GET("/", services.FindAllUsers)
	usersRoute.GET("/:id", services.FindByUserId)
	usersRoute.POST("/", services.CreateUser)
	usersRoute.DELETE("/:id", services.DeleteUser)
	usersRoute.PUT("/:id", services.UpdateUser)
}
