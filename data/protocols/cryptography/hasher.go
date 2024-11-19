package cryptography

type Hasher interface {
	Hash(plainText string) (string, error)
}
