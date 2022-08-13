package personal

import (
	"crypto/sha256"
	"crypto/subtle"
)

// TODO Inject database connection and get user from there?

type Authenticator struct{}

func (s *Authenticator) Authenticate(username, password string) bool {
	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))
	expectedUsernameHash := sha256.Sum256([]byte("foo"))
	expectedPasswordHash := sha256.Sum256([]byte("bar"))
	usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:])
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:])
	return usernameMatch == 1 && passwordMatch == 1
}
