# vtuber-searcher

## 開發紀錄

### 11/3-index.html設計完成 

### 11/4-main.go基本架構完成 資料庫也建立完成

### 11/5-初代完成 感謝 @KoukeFoxes 的幫忙除錯


## 開發過程

### 1.寫好index.html 跟style.css

我也只是弄了一點點，沒整什麼大學問

### 2.golang gin架構設計

超級頭痛 完全0基礎 這是我的開發步驟

#### 第零步:設定環境

待補

#### 第一步:安裝gin

 ```go get -u github.com/gin-gonic/gin```
  
#### 第二步:寫個雛形的main.go
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
  
  #### 如果出現`no required module provides package github.com/gin-gonic/gin: `
  
  在跟main.go同目錄的地方建個空白檔go.mod
  
  然後cmd輸入
  
  ```go mod edit -module=(Project name) ```
  
  把(Project name)改成你的專案名稱
  
  接下來用VSCODE編輯go.mod你會發現module上面有一個灰灰的"Run go mod tidy"
  
  go.mod go.sum就完成了 我到這邊就沒有報錯了
  
  #### 每次Run code都會跳出一個很煩的防火牆視窗
  
  可以利用指令打包執行檔
  ```
  go build -o main.exe main.go
	
  ```
  不一定要取main.exe 可以自己改 這樣開啟舒服多了
  
  
  #### 第三步:無情的撰寫功能
  
  ## 支援清單
  
  待補
