package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

// GenerateKey generates new RSA private key
func GenerateKey() (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println("unable to generate new rsa key: ", err)
		return nil, nil
	}
	return key, nil
}
