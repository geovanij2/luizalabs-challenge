package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"
)

type UpdateClient struct {
	clientRepository repository.ClientRepository
}

func NewUpdateClient(clientRepository repository.ClientRepository) *UpdateClient {
	return &UpdateClient{
		clientRepository: clientRepository,
	}
}

func (u *UpdateClient) Execute(client *entity.Client) error {
	existingClient, err := u.clientRepository.FindById(client.Id)

	if err != nil {
		return err
	}

	if existingClient == nil {
		return ErrClientNotFound
	}

	if existingClient.Email != client.Email {
		userWithEmail, err := u.clientRepository.FindByEmail(client.Email)
		if err != nil {
			return err
		}

		if userWithEmail != nil {
			return ErrClientEmailAlreadyExists
		}
	}

	_, err = u.clientRepository.Update(client)

	if err != nil {
		return err
	}

	return nil
}
