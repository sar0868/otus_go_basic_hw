package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/sar0868/otus_go_basic_hw/hw15_go_sql/docs"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title HW15: shop
// @version 1
// @description API Server

// @host 127.0.0.1:8080/

func main() {
	var IP string
	var PORT string
	flag.StringVar(&IP, "ip", "127.0.0.1", "- ip for server")
	flag.StringVar(&PORT, "port", "8080", "- port for server")
	flag.Parse()

	router := gin.Default()

	handle := handler.New()

	PATH := fmt.Sprintf("%s:%s", IP, PORT)
	url := ginSwagger.URL("http://" + PATH + "/swagger/doc.json")
	// ginSwagger.WrapHandler(swaggerFiles.Handler,
	// 	ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	// 	ginSwagger.DefaultModelsExpandDepth(-1))
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	handler.InitHandler(router, handle)

	if err := router.Run(PATH); err != nil {
		panic(err)
	}
}
