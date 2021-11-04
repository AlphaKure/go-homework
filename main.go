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
			"error":"json error",

		})
	}
	var data []JSONType
	json.Unmarshal([]byte(file), &data)
	var i int
	check := 11
	for i = 0; i < len(data); i++ {
		if i_id==data[i].Id {
			check = i
		}
	}
	if check == 11 {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error":"找不到資料",
			"i_id":i_id,
		})
		return
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"success": "找到了",
			"name":data[check].Name,
			"birthday":data[check].Birthday,
			"YT":data[check].YT,
			"twitter":data[check].Twitter,
			"i_id":i_id,
		})
		return
	}
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("html/*.html")
	server.Static("/style", "./html/style")
	server.GET("/", index)
	server.POST("/result", search)
	server.Run(":5501")
}
