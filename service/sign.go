package service

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"log"
)

func Sign(hexPrivateKey string, docHash []byte) (string, error) {

	private_key, err := DecodePrivate(hexPrivateKey)
	if err != nil {
		return "", err
	}

	signature, err := ecdsa.SignASN1(rand.Reader, private_key, docHash[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil

}

func Verify(signature string, publicKey string, docHash []byte) (bool, error) {
	byteSignature, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	PubKey, err := DecodePublic(publicKey)
	if err != nil {
		return false, err
	}

	check := ecdsa.VerifyASN1(PubKey, docHash[:], byteSignature)
	log.Println(check)
	return check, nil
}
