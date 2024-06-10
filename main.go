package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	fmt.Println("Hello World")

	r := router()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// 関数を外出しする手法
	r.GET("/quickstart", quickStart)

	// ミドルウェアの利用
	ua := ""
	r.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
		})
	})

	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}

func quickStart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "quickStart",
	})
}

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/ps", func(c *gin.Context) {
		var u User
		if err := c.BindJSON(&u); err != nil {
			// const StatusUnauthorized untyped int = 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": u})
	})

	return r
}
