package services

import (
	"documentation/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userList = []models.User{}

// Get User
func FindAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userList)
}

// Post User
func PostUser(ctx *gin.Context) {
	user := models.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	userList = append(userList, user)
	ctx.JSON(http.StatusOK, "Successfully posted")
}

// Delete User
func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	for _, user := range userList {
		log.Println(user)
		userList = append(userList[:userId], userList[userId+1:]...)
		ctx.JSON(http.StatusOK, "Successfully deleted")
		return
	}
	ctx.JSON(http.StatusNotFound, "Error")
}

// Put User
func PutUser(ctx *gin.Context) {
	updatedUser := models.User{}
	err := ctx.BindJSON(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
		return
	}
	userId, _ := strconv.Atoi(ctx.Param("id"))
	for key, user := range userList {
		if user.Id == userId {
			userList[key] = updatedUser
			log.Println(userList[key])
			ctx.JSON(http.StatusOK, "Successfully updated")
			return
		}
	}
	ctx.JSON(http.StatusNotFound, "Error")
}
