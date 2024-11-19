package application

import (
	"luizalabs-chalenge/data/protocols/cryptography"
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"
)

type Authentication struct {
	clientRepository repository.ClientRepository
	decrypter        cryptography.Decrypter
}

func NewAuthentication(clientRepository repository.ClientRepository, decrypter cryptography.Decrypter) *Authentication {
	return &Authentication{
		clientRepository: clientRepository,
		decrypter:        decrypter,
	}
}

func (a *Authentication) Execute(accessToken string) (*entity.Client, error) {
	clientId, err := a.decrypter.Decrypt(accessToken)

	if err != nil {
		return nil, err
	}

	client, err := a.clientRepository.FindById(clientId)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, ErrClientNotFound
	}

	return client, nil
}
