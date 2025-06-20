package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomSequence(n int) string {
	b := make([]rune, n)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range b {
		b[i] = letters[r1.Intn(999999)%len(letters)]
	}
	return string(b)
}

// GenSalt is used to combine with password, generate a different hash for the same password
func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randomSequence(length)
}
