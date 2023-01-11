package service

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func HashPDFfile(url string) []byte {
	file, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		panic(err)
	}
	return hash.Sum(nil)
}
