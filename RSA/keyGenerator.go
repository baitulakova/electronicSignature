package RSA

import (
	"crypto/rsa"
	"fmt"
	"crypto/rand"
)

func GenerateKey() *rsa.PrivateKey{
	PrKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error in key generation: ",err)
	}
	return PrKey
}
