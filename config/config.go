package config

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ContractAddress common.Address

var Client *ethclient.Client

func ConnectClient() *ethclient.Client {
	client, err := ethclient.Dial("https://rpc-mumbai.maticvigil.com/")
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	log.Printf("%s", client)
	return Client
}

func GetClient() *ethclient.Client {
	return Client
}
