package service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"log"
)

type EllipticCurve struct {
	pubKeyCurve elliptic.Curve
	privateKey  *ecdsa.PrivateKey
	publicKey   *ecdsa.PublicKey
}

func New(curve elliptic.Curve) *EllipticCurve {
	return &EllipticCurve{
		pubKeyCurve: curve,
		privateKey:  new(ecdsa.PrivateKey),
	}
}

func (ec *EllipticCurve) GenerateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {

	var err error
	privKey, err := ecdsa.GenerateKey(ec.pubKeyCurve, rand.Reader)

	if err == nil {
		ec.privateKey = privKey
		ec.publicKey = &privKey.PublicKey
	}
	log.Println(ec.privateKey)
	return ec.privateKey, ec.publicKey, err
}

func EncodePrivate(privKey *ecdsa.PrivateKey) (string, error) {
	encoded, err := x509.MarshalECPrivateKey(privKey)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encoded), nil
}

func EncodePublic(pubKey *ecdsa.PublicKey) string {
	if pubKey == nil || pubKey.X == nil || pubKey.Y == nil {
		return ""
	}

	encoded := elliptic.Marshal(elliptic.P256(), pubKey.X, pubKey.Y)
	return hex.EncodeToString(encoded)
}

func DecodePrivate(hexPrvKey string) (*ecdsa.PrivateKey, error) {
	var privKey *ecdsa.PrivateKey

	bytesPrvKey, err := hex.DecodeString(hexPrvKey)
	if err != nil {
		return privKey, err
	}

	privateKey, err := x509.ParseECPrivateKey(bytesPrvKey)
	if err != nil {
		return privKey, err
	}
	log.Println(privateKey)
	return privateKey, err
}

func DecodePublic(hexPubKey string) (*ecdsa.PublicKey, error) {
	var pubKey *ecdsa.PublicKey

	bytesPubKey, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return pubKey, err
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), bytesPubKey)

	return &ecdsa.PublicKey{
		elliptic.P256(),
		x, y,
	}, nil

}
