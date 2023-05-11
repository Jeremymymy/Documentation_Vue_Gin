package services

import (
	"documentation/middlewares"
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

// Get User by Email
func FindByUserEmail(ctx *gin.Context) {
	user := models.FindByUserEmail(ctx.Param("email"))
	if user.EmployeeId == "" {
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
	if user.EmployeeId == "" {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Login User
func LoginUser(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	isMatchedPass := models.CheckUserPassword(email, password)
	if isMatchedPass == false {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}
	user := models.FindByUserEmail(email)
	middlewares.SaveSession(ctx, user.EmployeeId)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Login Successfully",
		"User":     user,
		"Sessions": middlewares.GetSession(ctx),
	})
}

// Logout User
func LogoutUser(ctx *gin.Context) {
	middlewares.ClearSession(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout Successfully",
	})
}

// Check User Session
func CheckUserSession(ctx *gin.Context) {
	sessionId := middlewares.GetSession(ctx)
	if sessionId == "0" {
		ctx.JSON(http.StatusUnauthorized, "Error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Check Session Successfully",
		"User":    sessionId,
	})
}
