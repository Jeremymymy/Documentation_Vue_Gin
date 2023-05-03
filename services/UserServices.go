package services

import (
	"documentation/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User
func FindAllUsers(ctx *gin.Context) {
	users := models.FindAllUsers()
	ctx.JSON(http.StatusOK, users)
}

// Get User by Id
func FindByUserId(ctx *gin.Context) {
	user := models.FindByUserId(ctx.Param("id"))
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	ctx.JSON(http.StatusOK, user)
}

// Create User
func CreateUser(ctx *gin.Context) {
	user := models.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	newUser := models.CreateUser(user)
	ctx.JSON(http.StatusOK, newUser)
}

// Delete User
func DeleteUser(ctx *gin.Context) {
	isDeleted := models.DeleteUser(ctx.Param("id"))
	if !isDeleted {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	ctx.JSON(http.StatusOK, "Successfully")
}

// Update User
func UpdateUser(ctx *gin.Context) {
	user := models.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = models.UpdateUser(ctx.Param("id"), user)
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	ctx.JSON(http.StatusOK, user)
}
