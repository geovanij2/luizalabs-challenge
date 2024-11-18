package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"

	"github.com/google/uuid"
)

type CreateClient struct {
	clientRepository repository.ClientRepository
}

type CreateClientInput struct {
	Email    string
	Name     string
	Password string
}

func NewCreateClient(repository repository.ClientRepository) *CreateClient {
	return &CreateClient{
		clientRepository: repository,
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

	newClient := entity.Client{
		Id:       uuid.New().String(),
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password, // TODO salvar hash da senha
	}

	client, err := c.clientRepository.Create(&newClient)

	if err != nil {
		return nil, err
	}

	return client, nil
}
