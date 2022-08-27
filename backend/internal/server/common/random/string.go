package random

import "math/rand"

// String builds a new random alphanumerical string with the given length n.
func String(n int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s := make([]byte, n)
	for i := range s {
		//nolint:gosec
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}
