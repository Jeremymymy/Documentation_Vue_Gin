package services

import (
	"documentation/middlewares"
	"documentation/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Document and Create Doc_version1
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
	doc.Belong = user.Department
	newDoc, _ := models.CreateDoc(doc)
	if err := models.AddNewDoc(user, newDoc); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}

	ver := models.Version{}
	ver.DocId = newDoc.ID
	ver.UpdaterId = user.EmployeeId
	ver.UpdaterEmail = user.Email
	ver.UpdaterName = user.Name
	ver.Title = newDoc.Title
	ver.Content = newDoc.Content
	newVer, _ := models.CreateVer(ver)
	if err := models.DocAddNewVer(newDoc, newVer); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Document created successfully",
		"newDoc":  newDoc,
	})
}

// Get All stored Versions
func GetAllVers(ctx *gin.Context) {
	vers := models.GetAllVers()
	ctx.JSON(http.StatusOK, vers)
}

// Get Document by DocId
func GetDocById(ctx *gin.Context) {
	docId, _ := strconv.Atoi(ctx.Param("docId"))
	doc, err := models.GetDocById(uint(docId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Error : "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, doc)
}

// Get Document and its All Versions by DocId
func GetDocByIdWithVersPreload(ctx *gin.Context) {
	docId, _ := strconv.Atoi(ctx.Param("docId"))
	doc, err := models.GetDocByIdWithVersPreload(uint(docId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Error : "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, doc)
}

// Delete Document by DocId
func DeleteDoc(ctx *gin.Context) {
	docId, _ := strconv.Atoi(ctx.Param("docId"))
	isDeleted := models.DeleteDoc(uint(docId))
	if !isDeleted {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Fail to delete this doc"})
		return
	}

	ctx.JSON(http.StatusOK, "Successfully deleted")
}

// Update Document and Add new Doc_version by DocId
func UpdateDoc(ctx *gin.Context) {
	sessionId := middlewares.GetSession(ctx)
	if sessionId == "0" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	user, _ := models.FindByUserIdWithEditVersPreload(sessionId)
	if user.EmployeeId == "" {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}

	docId, _ := strconv.Atoi(ctx.Param("docId"))
	doc, err := models.GetDocByIdWithVersPreload(uint(docId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Error : "+err.Error())
		return
	}
	if err := ctx.BindJSON(&doc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document data"})
		return
	}

	ver := models.Version{}
	ver.DocId = doc.ID
	ver.VerNum = len(doc.Vers) + 1
	ver.UpdaterId = user.EmployeeId
	ver.UpdaterEmail = user.Email
	ver.UpdaterName = user.Name
	ver.Title = doc.Title
	ver.Content = doc.Content
	newVer, _ := models.CreateVer(ver)
	if err := models.DocAddNewVer(doc, newVer); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}
	if err := models.UserAddNewVer(user, newVer); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}

	updatedDoc, err := models.UpdateDoc(uint(docId), doc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"updatedDoc": updatedDoc})
}

// Add Collection
func CollectDoc(ctx *gin.Context) {
	sessionId := middlewares.GetSession(ctx)
	if sessionId == "0" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	user, _ := models.FindByUserIdWithCollectDocsPreload(sessionId)
	if user.EmployeeId == "" {
		ctx.JSON(http.StatusNotFound, "Error")
		return
	}

	docId, _ := strconv.Atoi(ctx.Param("docId"))
	doc, err := models.GetDocById(uint(docId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Error : "+err.Error())
		return
	}
	collect := models.Collect{}
	collect.UserId = user.EmployeeId
	collect.AuthorId = doc.AuthorId
	collect.AuthorName = doc.AuthorName
	collect.Belong = doc.Belong
	collect.Title = doc.Title
	collect.Content = doc.Content
	col, _ := models.CreateCollect(collect)
	if err := models.UserAddCollect(user, col); err != nil {
		ctx.JSON(http.StatusInternalServerError, "Error : "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Collect": col,
		"User":    user,
	})
}

// Get All stored Collections
func GetAllCollects(ctx *gin.Context) {
	cols := models.GetAllCollects()
	ctx.JSON(http.StatusOK, cols)
}

// Delete Collection
func DeleteCol(ctx *gin.Context) {
	colId, _ := strconv.Atoi(ctx.Param("colId"))
	isDeleted := models.DeleteCol(uint(colId))
	if !isDeleted {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Fail to delete this collect"})
		return
	}
	ctx.JSON(http.StatusOK, "Successfully deleted")
}
