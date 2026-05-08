package auth_infra

import "golang.org/x/crypto/bcrypt"

type AuthHashing struct{}

func NewAuthHashing() *AuthHashing {
	return &AuthHashing{}
}

func (a *AuthHashing) Hashing(s string) ([]byte, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	return hashPassword, err
}

func (a *AuthHashing) CheckPassword(hashedPassword, userEnteredPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userEnteredPassword))
	return err
}
