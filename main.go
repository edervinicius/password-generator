package main

import (
	"fmt"
	"math/big"
	"crypto/rand"
)

func generateRandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?`~"
	charsetLength := big.NewInt(int64(len(charset)))

	password := make([]byte, length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}

	return string(password), nil
}

func main() {
	passwordLength := 12
	password, err := generateRandomPassword(passwordLength)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(password)
}