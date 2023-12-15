package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin의 기본 라우터를 생성
	router := gin.Default()

	// 루트 엔드포인트
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	// /ping 엔드포인트
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// /hello/:name 엔드포인트 (파라미터 전달)
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
	})

	// 서버 시작
	router.Run(":8080")
}
