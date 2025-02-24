package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/sar0868/otus_go_basic_hw/hw15_go_sql/docs"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title HW15: shop
// @version 1
// @description API Server

// @host 127.0.0.1:8080/

func main() {
	conf := app.Init()

	IP := conf.HTTP.Host
	PORT := conf.HTTP.Port

	router := gin.Default()

	handle := handler.New()

	PATH := fmt.Sprintf("%s:%d", IP, PORT)
	url := ginSwagger.URL("http://" + PATH + "/swagger/doc.json")
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	handler.InitHandler(router, handle)

	if errRouter := router.Run(PATH); errRouter != nil {
		log.Panicln(errRouter.Error())
	}
}
