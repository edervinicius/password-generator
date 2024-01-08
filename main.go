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
	var passwordLength int

	fmt.Print("Input the password length: ")
	fmt.Scan(&passwordLength)

	password, err := generateRandomPassword(passwordLength)

	if err != nil {
		fmt.Println("Erro ao gerar senha:", err)
		return
	}

	fmt.Println("Password:", password)
}