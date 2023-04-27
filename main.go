package main

import (
	. "documentation/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	home := router.Group("/")
	AddUserRouter(home)
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "root",
	// 	})
	// })

	// router.POST("/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	ctx.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })

	router.Run(":8000")
}
