package core

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randSequence
func randSequence(n int) string {
	b := make([]rune, n)

	s1 := rand.NewSource(time.Now().UnixNano())

	r1 := rand.New(s1)

	lettersLen := len(letters)

	for i := range b {
		b[i] = letters[r1.Intn(999999)%lettersLen]
	}

	return string(b)
}

// GenSalt function generates a random string of a specified length using the randSequence function. If the provided length is less than 0, it defaults to a length of 50. This function is typically used to generate a salt for cryptographic operations
func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}
