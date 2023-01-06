package controller

import (
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dattvan/plg-doc-vertify/api"
	"github.com/dattvan/plg-doc-vertify/config"
	"github.com/dattvan/plg-doc-vertify/models"
	"github.com/dattvan/plg-doc-vertify/service"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StorageContract(c *gin.Context) {
	godotenv.Load(".env")
	arg := models.StoreContractArgument{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	add := common.HexToAddress("0x7f4a627a135cb185f61e00deb24c213c394be0ff")
	client := config.GetClient()
	instance, err := api.NewApi(add, client)
	if err != nil {
		log.Fatal(err)
	}

	abi, err := abi.JSON(strings.NewReader(models.AbiDefinition))
	if err != nil {
		panic(err)
	}
	out, err := abi.Pack("addContract", string(arg.Name), string(arg.Company), string(arg.IdentityCard), string(arg.HashDocs))
	if err != nil {
		panic(err)
	}

	auth := service.GetAccountAuth(client, os.Getenv("PRV_KEY"), out)

	tx, err := instance.AddContract(auth, arg.Name, arg.Company, arg.IdentityCard, arg.HashDocs)
	if err != nil {
		log.Fatal(err)
	}

	txResponse := models.Transaction{
		Hash:     tx.Hash().Hex(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		Data:     hex.EncodeToString(tx.Data()),
		GasPrice: tx.GasPrice().Uint64(),
	}
	c.JSON(http.StatusCreated, gin.H{"transaction": txResponse})
}

func GetContract(c *gin.Context) {
	godotenv.Load(".env")
	arg := models.GetContractArgument{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	add := common.HexToAddress("0x7f4a627a135cb185f61e00deb24c213c394be0ff")

	client := config.GetClient()
	instance, err := api.NewApi(add, client)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	tx, err := instance.GetContract(nil, arg.IdentityCard, arg.Company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contractResponse := models.GetContractResponse{
		HashDoc: tx,
	}
	c.JSON(http.StatusOK, gin.H{"hash_doc": contractResponse})

}
