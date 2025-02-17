package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title HW15: shop
// @version 1
// @description API Server

// @host localhost:8080/

func main() {
	router := gin.Default()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	router.Run(":8080")
}
