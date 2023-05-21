package services

import (
	"documentation/middlewares"
	"documentation/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDoc(ctx *gin.Context) {
	sessionId := middlewares.GetSession(ctx)
	if sessionId == "0" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	user, _ := models.FindByUserIdWithPostDocsPreload(sessionId)

	doc := models.Document{}
	if err := ctx.BindJSON(&doc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document data"})
		return
	}
	doc.AuthorId = sessionId
	doc.AuthorName = user.Name
	newDoc, _ := models.CreateDoc(doc)
	if err := models.AddNewDoc(user, newDoc); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Document created successfully",
		"newDoc":  newDoc,
	})
}
