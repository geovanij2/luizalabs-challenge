package cryptography

type HashComparer interface {
	Compare(plainText, hashedText string) (bool, error)
}
