package routers

import (
	session "documentation/middlewares"
	"documentation/services"

	"github.com/gin-gonic/gin"
)

func AddDocRouter(rg *gin.RouterGroup) {
	docsRoute := rg.Group("/docs", session.SetSession())

	docsRoute.GET("/getPublicDocs/", services.GetPublicDocs)

	docsRoute.Use(session.AuthSession())
	{
		docsRoute.POST("/createDoc", services.CreateDoc)
		docsRoute.GET("/getDoc/:docId", services.GetDocById)
		docsRoute.GET("/getDepartmentDocs/:department", services.GetDepartmentDocs)
		docsRoute.GET("/getDocAllVers/:docId", services.GetDocByIdWithVersPreload)
		docsRoute.GET("/getAllVers/", services.GetAllVers)
		docsRoute.GET("/getDocVer/:verId", services.GetDocVer)
		docsRoute.GET("/collectDoc/:docId", services.CollectDoc)
		docsRoute.GET("/getAllCollects/", services.GetAllCollects)
		docsRoute.DELETE("/deleteDoc/:docId", services.DeleteDoc)
		docsRoute.DELETE("/deleteCollect/:colId", services.DeleteCol)
		docsRoute.PUT("/update/:docId", services.UpdateDoc)
	}
}
