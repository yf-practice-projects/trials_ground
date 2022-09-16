package main

import (
	"html/template"
	"time"
	"web/trials_ground/routes"
	"web/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	time.LoadLocation("Asia/Tokyo")
	router := routes.NewRoutes()
	router.SetFuncMap(template.FuncMap{
		"convertTimeToDateTime": utils.ConvertTimeToDateTime,
	})
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })
	// router.GET("/test", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index copy.html", gin.H{})
	// })

	router.Run()
}
