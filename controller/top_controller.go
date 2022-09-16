package controller

import (
	"net/http"
	"web/trials_ground/models"

	"github.com/gin-gonic/gin"
)

func AllArticlesView(c *gin.Context) {
	//c.JSON(http.StatusOK, articles)
	articlesData := models.GetAll()
	c.HTML(http.StatusOK, "topPage.html", gin.H{"datas": articlesData})
}

func OneArticleView(c *gin.Context) {
	id := c.Param("id")
	article := models.GetArticle(id)
	c.HTML(http.StatusOK, "top.html", gin.H{"article": article})
}

func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "index copy.html", gin.H{})
}
