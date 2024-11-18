package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"
)

type ReadClient struct {
	clientRepository repository.ClientRepository
}

type ReadClientInput struct {
	ClientId string
}

func NewReadClient(clientRepository repository.ClientRepository) *ReadClient {
	return &ReadClient{
		clientRepository: clientRepository,
	}
}

func (r *ReadClient) Execute(input ReadClientInput) (*entity.Client, error) {
	client, err := r.clientRepository.FindById(input.ClientId)

	if err != nil {
		return nil, err
	}

	if client == nil {
		return nil, ErrClientNotFound
	}

	return client, nil
}
