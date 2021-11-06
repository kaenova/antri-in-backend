package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	var hash string
	hashByte, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("cannot hash password")
	}
	hash = string(hashByte)
	return hash, nil
}

func CompareHashPassword(input string, compared string) bool {
	inputByte := []byte(input)
	comparedByte := []byte(compared)
	err := bcrypt.CompareHashAndPassword(comparedByte, inputByte)
	return err == nil
}
