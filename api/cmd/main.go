package main

import (
	"fmt"

	"event-driven/config"
	"event-driven/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	router := gin.Default()
	router.POST("/messages", handler.HandlePostMessage)
	fmt.Printf("config.CFG.ApiPort: %s\n", config.CFG.ApiPort)
	runString := fmt.Sprintf(":%s", config.CFG.ApiPort)
	router.Run(runString)
}
