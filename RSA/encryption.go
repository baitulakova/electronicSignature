package RSA

import (
	"crypto/rsa"
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/rand"
	"crypto"
	"hash"
)

func Encrypt(targetPubKey *rsa.PublicKey,text []byte)(ciphertext,label []byte,randomHash hash.Hash){
	label=[]byte("")
	randomHash=sha256.New()
	ciphertext, err := rsa.EncryptOAEP(randomHash, rand.Reader, targetPubKey, text, label)
	if err != nil {
		fmt.Println("Error in encryption: ",err)
		os.Exit(1)
	}
	return ciphertext,label,randomHash
}

func SignMessage(text string,privKey *rsa.PrivateKey)(signature []byte,hashed []byte,Opts rsa.PSSOptions){
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	newhash := crypto.SHA256 
	Hash := newhash.New()
	Hash.Write([]byte(text))
	hashed = Hash.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, privKey, newhash, hashed, &opts)
	if err != nil {
		fmt.Println("Error in signing message: ", err)
		os.Exit(1)
	}
	return signature,hashed,opts
}

func Decrypt(cipherText []byte,targetPrivKey *rsa.PrivateKey,randomHash hash.Hash,label []byte)(plainText string){
	pText, err := rsa.DecryptOAEP(randomHash, rand.Reader, targetPrivKey, cipherText, label)
	if err != nil {
		fmt.Println("Error in decryption: ", err)
		os.Exit(1)
	}
	fmt.Println("Plain text: ", string(pText))
	plainText=string(pText)
	return plainText
}

func Verify(PublicKey *rsa.PublicKey,hashed,signature []byte,opts rsa.PSSOptions){
	newhash := crypto.SHA256
	err := rsa.VerifyPSS(PublicKey, newhash, hashed, signature, &opts)
	if err != nil {
		fmt.Println("Verify signature failed")
		os.Exit(1)
	} else {
		fmt.Println("Verify signature successful")
	}
}
