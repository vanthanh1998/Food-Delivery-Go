package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string { // private
	b := make([]rune, n)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range b {
		b[i] = letters[r1.Intn(99999)%len(letters)]
	}
	return string(b)
}
func GenSalt(length int) string { // public
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}
