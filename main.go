package main

import (
	"documentation/dbconnect"
	src "documentation/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	home := router.Group("/TSMC")
	src.AddUserRouter(home)

	go func() {
		dbconnect.DD()
	}()

	router.Run(":8000")
}
