package cryptography

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct{}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{}
}

func (b *BcryptAdapter) Hash(plainText string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (b *BcryptAdapter) Compare(plainText, hashedText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))

	if err != nil {
		return false, err
	}

	return true, nil
}
