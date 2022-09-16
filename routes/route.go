package routes

import (
	"web/trials_ground/controller"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()

	// 記事一覧
	router.GET("/", controller.AllArticlesView)
	router.GET("/golang/:id", controller.OneArticleView)
	router.GET("/test", controller.Test)

	return router

}
