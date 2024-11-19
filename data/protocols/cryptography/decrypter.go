package cryptography

type Decrypter interface {
	Decrypt(encryptedText string) (string, error)
}
