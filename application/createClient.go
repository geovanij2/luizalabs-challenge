package application

import (
	"luizalabs-challenge/data/protocols/cryptography"
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/domain/repository"

	"github.com/google/uuid"
)

type CreateClient struct {
	clientRepository repository.ClientRepository
	hasher           cryptography.Hasher
}

type CreateClientInput struct {
	Email    string
	Name     string
	Password string
}

func NewCreateClient(repository repository.ClientRepository, hasher cryptography.Hasher) *CreateClient {
	return &CreateClient{
		clientRepository: repository,
		hasher:           hasher,
	}
}

func (c *CreateClient) Execute(input CreateClientInput) (*entity.Client, error) {
	existingUser, err := c.clientRepository.FindByEmail(input.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrClientEmailAlreadyExists
	}

	passwordHash, err := c.hasher.Hash(input.Password)

	if err != nil {
		return nil, err
	}

	newClient := entity.Client{
		Id:       uuid.New().String(),
		Email:    input.Email,
		Name:     input.Name,
		Password: passwordHash,
	}

	client, err := c.clientRepository.Create(&newClient)

	if err != nil {
		return nil, err
	}

	return client, nil
}
