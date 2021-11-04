package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html",nil)
}

func result(c *gin.Context) {
	c.HTML(http.StatusOK, "result.html", nil)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("html/*.html")
	server.Static("/style", "./html/style")
	server.GET("/", index)
	server.POST("/result",nil)
	server.GET("/result", result)
	server.Run(":5501")
}
