package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordBcrypt(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", hash), nil
}

func ComparePasswords(hashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
