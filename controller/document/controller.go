package document

import (
	"crypto/elliptic"
	"log"
	"net/http"

	"github.com/dattvan/plg-doc-vertify/models"
	"github.com/dattvan/plg-doc-vertify/service"
	"github.com/gin-gonic/gin"
)

func SignContract(c *gin.Context) {
	arg := models.SignContractArgument{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	contractBytes := service.HashPDFfile(arg.Url)
	signature, err := service.Sign(arg.PrivateKey, contractBytes)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	response := &models.SignatureResponse{
		Signature: signature,
	}

	c.JSON(http.StatusOK, gin.H{"response": response})

}

func VerifyContract(c *gin.Context) {
	arg := models.VerifyContractArgument{}
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	contractBytes := service.HashPDFfile(arg.Url)
	verify, err := service.Verify(arg.Signature, arg.PublicKey, contractBytes)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	log.Println(verify)
	if !verify {
		c.JSON(http.StatusOK, gin.H{"message": "document not validated"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "document validated"})
	}

}

func GenerateKeyPair(c *gin.Context) {
	eCKey := service.New(elliptic.P256())
	privateKey, publicKey, err := eCKey.GenerateKeys()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	encodePrvKey, err := service.EncodePrivate(privateKey)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	encodePubKey := service.EncodePublic(publicKey)

	response := models.KeyPair{
		PrivateKey: encodePrvKey,
		PublicKey:  encodePubKey,
	}
	c.JSON(http.StatusOK, gin.H{"response": response})

}
