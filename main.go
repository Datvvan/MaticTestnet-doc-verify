package main

import (
	"github.com/dattvan/plg-doc-vertify/config"
	"github.com/dattvan/plg-doc-vertify/controller"
	"github.com/dattvan/plg-doc-vertify/controller/document"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectClient()
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/storagecontract", controller.StorageContract)
		api.GET("/getcontract", controller.GetContract)
		api.GET("/genkey", document.GenerateKeyPair)
		api.POST("/signature", document.SignContract)
		api.POST("/verify", document.VerifyContract)

	}
	router.Run(":8080")
}
