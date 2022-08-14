package admin

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/toivjon/max-pondus/backend/internal/server/common"
)

// TODO Inject database connection and get user from there?

type Authenticator struct{}

func (s *Authenticator) Authenticate(username, password string) (bool, common.User) {
	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))
	expectedUsernameHash := sha256.Sum256([]byte("foo"))
	expectedPasswordHash := sha256.Sum256([]byte("bar"))
	usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:])
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:])
	if usernameMatch == 1 && passwordMatch == 1 {
		return true, common.User{
			Username: username,
		}
	}
	return false, common.User{}
}
