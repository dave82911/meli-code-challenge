package router

import (
	"meli-code-challenge/controller"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.ItemsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/meli-challenge")
	tagsRouter.POST("/process-items", tagsController.ProcessItemsFile)

	return router
}
