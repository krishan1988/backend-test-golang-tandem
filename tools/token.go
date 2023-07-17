// Package main (token.go) is responsible for generation tokens for testing purpose.
package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	var signKey *rsa.PrivateKey
	signBytes, err := ioutil.ReadFile("private.key")
	if err != nil {
		log.Fatal(err.Error())
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
		Issuer:    "",
	}
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(tokenString)
}
