package cryptography

type Encrypter interface {
	Encrypt(plainText string) (string, error)
}
