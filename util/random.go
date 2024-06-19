package util

import "math/rand"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func RandomString(length int) string {
	randomString := ""

	for i := 0; i < length; i++ {
		randomAlphabetIndex := rand.Intn(len(alphabet))
		randomString += string(alphabet[randomAlphabetIndex])
	}

	return randomString
}

func RandomUsername() string {
	nameLength := rand.Intn(10) + 2
	return RandomString(nameLength)
}
