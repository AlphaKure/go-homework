package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func result(c *gin.Context) {
	c.HTML(http.StatusOK, "result.html", nil)
}

type JSONType struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	YT       string `json:"YT"`
	Twitter  string `json:"twitter"`
}

func search(c *gin.Context) {
	i_id, _ := c.GetPostForm("id_input")
	file, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "json error",
		})
	}
	var data []JSONType
	json.Unmarshal([]byte(file), &data)
	var i int
	for i = 0; i < len(data); i++ {
		if i_id == data[i].Id {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"btn":      "result",
				"success":  "finded",
				"Title":    "結果",
				"img":      "/style/pic/" + data[i].Id + ".jpg",
				"name":     data[i].Name,
				"birthday": data[i].Birthday,
				"YT":       data[i].YT,
				"twitter":  data[i].Twitter,
			})
			return
		}
	}
	c.HTML(http.StatusBadRequest, "index.html", gin.H{
		"btn":   "result",
		"Title": "結果",
		"error": "找不到資料",
	})
	return
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("html/*.html")
	server.Static("/style", "./html/style")
	server.GET("/", index)
	server.POST("/result", search)
	server.Run(":5501")
}
