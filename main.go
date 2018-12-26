package main

import (
	"net/http"
	"log"
	"crypto/rsa"
	"github.com/baitulakova/Electronic_signature_RSA/vars"
	"github.com/baitulakova/Electronic_signature_RSA/RSA"
	"html/template"
)

type Member struct {
	FirstName string
	LastName string
	PublicKey *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

var Members=make(map[string]Member)

func registerHandler(w http.ResponseWriter,r *http.Request){
	w.Write(vars.LoginPage)
}

func registerStatusHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fName:=r.Form["fname"]
	lName:=r.Form["lname"]

	member:=Member{
		FirstName:fName[0],
		LastName:lName[0],
	}
	privKey:=RSA.GenerateKey()
	member.privateKey=privKey
	member.PublicKey=&privKey.PublicKey

	Members[member.FirstName]=member

	t,err:=template.New("webPage").Parse(vars.RegisterAnswer)
	if err!=nil{
		log.Println("Error parsing html: ",err)
		return
	}
	data := struct {
		FName string
		LName string
		PubKey *rsa.PublicKey
	}{
		FName:member.FirstName,
		LName:member.LastName,
		PubKey:member.PublicKey,
	}
	err = t.Execute(w, data)
	if err!=nil{
		log.Println("Error: ",err)
	}
}

func cipherHandler(w http.ResponseWriter,r *http.Request){
	w.Write(vars.CipherPage)
}

func decryptHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	plainText:=r.Form["plainText"]
	targetName:=r.Form["targetName"]

	target:=Members[targetName[0]]
	text:=[]byte(plainText[0])

	role:="sender"
	s:=Members[role]
	
	signature,hashed,opt:=RSA.SignMessage(plainText[0],s.privateKey)
	cipherText,label,hash:=RSA.Encrypt(target.PublicKey,text)

	t,err:=template.New("webPage").Parse(vars.CipherText)
	if err!=nil{
		log.Println("Error parsing html: ",err)
		return
	}
	log.Println(string(cipherText))
	data:= struct {
		cText string
	}{
		cText:string(cipherText),
	}
	t.Execute(w,data)

	//decryption
	user:="rec"
	rec:=Members[user]
	
	pText:=RSA.Decrypt(cipherText,rec.privateKey,hash,label)
	log.Println("Decrypted text: ",pText)
	
	RSA.Verify(s.PublicKey,hashed,signature,opt)
}

func main(){
	http.HandleFunc("/decrypt",decryptHandler)
	http.HandleFunc("/cipher",cipherHandler)
	http.HandleFunc("/member",registerStatusHandler)
	http.HandleFunc("/register",registerHandler)
	log.Println("Server is working on port :8080")
	http.ListenAndServe(":8080", nil)
}
