package random

import (
	"crypto/rand"
	"math/big"
)

// String builds a new random alphanumerical string with the given length n.
func String(n int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, n)
	for i := range bytes {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		bytes[i] = charset[num.Int64()]
	}
	return string(bytes)
}
