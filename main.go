package main

import (
	"documentation/dbconnect"
	"documentation/models"
	src "documentation/routers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run(":8000")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}                   // 允许的域名或IP地址
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // 允许的HTTP方法
	config.AllowCredentials = true
	router.Use(cors.New(config))

	home := router.Group("/TSMC")
	src.AddUserRouter(home)
	src.AddDocRouter(home)

	db := dbconnect.MySQLcon
	err := db.AutoMigrate(&models.User{}, &models.Document{}, &models.Version{}, &models.Collect{})
	if err != nil {
		log.Fatal(err)
	}

	return router
}
