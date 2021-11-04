# vtuber-searcher

## 開發紀錄

### 11/3-index.html設計完成 result.html基礎弄好

## 開發過程

### 1.寫好index.html 跟style.css

我也只是弄了一點點，沒整什麼大學問

### 2.golang gin架構設計

超級頭痛 完全0基礎 補充一下助教沒說到(還是我沒注意到的)

第一步:安裝gin

 ```go get -u github.com/gin-gonic/gin```
  
第二步:寫個雛形的main.go
 ```go
  package main
  import(
	  "net/http"
	  "github.com/gin-gonic/gin"
  )

  func index(c *gin.Context) {
	  c.HTML(http.StatusOK, "index.html",nil)
  }

  func main() {
	  server := gin.Default()
	  server.LoadHTMLGlob("html/*.html")
	  server.Static("/style", "./html/style")
	  server.GET("/", index)
	  server.Run(":5501")
  }
```
  上面main要依照你寫的做修改
  
  第三步:無情的撰寫功能
  
  待補
