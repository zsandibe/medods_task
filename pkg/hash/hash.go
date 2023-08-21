package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Generate(token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating hash: %v\n", err)
		return "", err
	}
	return string(bytes), nil
}

func Compare(token string, hashTokent string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashTokent), []byte(token))
	if err != nil {
		return false
	}
	return true
}
