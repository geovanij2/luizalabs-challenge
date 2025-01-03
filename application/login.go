package application

import (
	"luizalabs-challenge/data/protocols/cryptography"
	"luizalabs-challenge/domain/repository"
)

type Login struct {
	clientRepository repository.ClientRepository
	encrypter        cryptography.Encrypter
	hashComparer     cryptography.HashComparer
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLogin(
	clientRepository repository.ClientRepository,
	encrypter cryptography.Encrypter,
	hashComparer cryptography.HashComparer,
) *Login {
	return &Login{
		clientRepository: clientRepository,
		encrypter:        encrypter,
		hashComparer:     hashComparer,
	}
}

func (l *Login) Execute(input LoginInput) (string, error) {
	client, err := l.clientRepository.FindByEmail(input.Email)

	if err != nil {
		return "", err
	}

	if client == nil {
		return "", ErrClientNotFound
	}

	ok, _ := l.hashComparer.Compare(input.Password, client.Password)

	if !ok {
		return "", ErrWrongPassword
	}

	accessToken, err := l.encrypter.Encrypt(client.Id)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
