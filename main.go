package main

import (
	"log"

	"github.com/dattvan/plg-doc-vertify/config"
	"github.com/dattvan/plg-doc-vertify/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	client := config.ConnectClient()
	log.Printf("%s", client)
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/storagecontract", controller.StorageContract)
		api.GET("/getcontract", controller.GetContract)

	}
	router.Run(":8080")
}
