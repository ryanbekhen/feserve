package cert

import (
	"crypto"

	"github.com/go-acme/lego/v4/registration"
)

// You'll need a user or account type that implements acme.User
type User struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *User) GetEmail() string {
	return u.Email
}
func (u User) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.key
}
