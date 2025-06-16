package main

import (
	"bin_blog/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var Router *gin.Engine

func main() {
	router := gin.Default()
	Router = router

	config.LoadConfig()

	fmt.Println(config.Global.App.Env)
	// 设置生产模式
	if config.Global.App.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}

	err := router.Run(fmt.Sprintf(
		"%s:%s",
		config.Global.App.Host,
		config.Global.App.Port,
	))

	if err != nil {
		log.Fatal(err)
	}
}
