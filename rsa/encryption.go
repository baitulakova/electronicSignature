package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

// Encrypt encrypts text message
func Encrypt(targetPubKey *rsa.PublicKey, text []byte) ([]byte, error) {
	// The label parameter may contain arbitrary data that will not be encrypted, but which gives important context to the message
	label := []byte("")
	// The random parameter is used as a source of entropy to ensure that encrypting the same message twice doesn't result in the same ciphertext
	randomHash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(randomHash, rand.Reader, targetPubKey, text, label)
	if err != nil {
		log.Print("unable to encrypt: ", err)
		return nil, err
	}
	return ciphertext, nil
}

// SignMessage calculates the signature of hashed using RSASSA-PSS
func SignMessage(text string, privKey *rsa.PrivateKey) ([]byte, error) {
	var opts rsa.PSSOptions

	opts.SaltLength = rsa.PSSSaltLengthAuto
	newhash := crypto.SHA256
	hash := newhash.New()
	hash.Write([]byte(text))
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privKey, newhash, hashed, &opts)
	if err != nil {
		log.Print("unable to signing message: ", err)
		return nil, err
	}
	return signature, nil
}

// Decrypt decrypts text message
func Decrypt(cipherText []byte, targetPrivKey *rsa.PrivateKey) (string, error) {
	var plainText string

	label := []byte("")
	randomHash := sha256.New()

	pText, err := rsa.DecryptOAEP(randomHash, rand.Reader, targetPrivKey, cipherText, label)
	if err != nil {
		log.Print("unable to decrypt text message: ", err)
		return plainText, err
	}
	plainText = string(pText)
	return plainText, nil
}

// Verify verifies signature
func Verify(pubKey *rsa.PublicKey, hashed, signature []byte) bool {
	newhash := crypto.SHA256
	opts := rsa.PSSOptions{}
	opts.SaltLength = rsa.PSSSaltLengthAuto

	err := rsa.VerifyPSS(pubKey, newhash, hashed, signature, &opts)
	if err != nil {
		log.Print("Verify signature failed")
		return false
	}
	return true
}
