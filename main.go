package main

import (
	"documentation/dbconnect"
	"documentation/models"
	src "documentation/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	home := router.Group("/TSMC")
	src.AddUserRouter(home)

	db := dbconnect.MySQLcon

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	router.Run(":8000")
}
