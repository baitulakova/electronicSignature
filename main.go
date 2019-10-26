package main

import (
	"fmt"

	"github.com/baitulakova/electronicSignature/rsa"
)

func main() {
	privKey, _ := rsa.GenerateKey()
	plainText := "hello gopher"

	// signature, _ := rsa.SignMessage(plainText, privKey)
	cipherText, _ := rsa.Encrypt(&privKey.PublicKey, []byte(plainText))

	fmt.Println(string(cipherText))

	pText, _ := rsa.Decrypt(cipherText, privKey)
	fmt.Println("Decrypted text: ", pText)

	// rsa.Verify(s.PublicKey, hashed, signature, opt)
}
