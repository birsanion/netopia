package rand

import (
	"math/rand"
)

const (
	LowerLetters  = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters       = LowerLetters + UpperLetters
	Digits        = "0123456789"
	DigitsLetters = Digits + Letters
)

func RandStringBytes(n int, alphabet string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}
