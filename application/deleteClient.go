package application

import "luizalabs-chalenge/domain/repository"

type DeleteClient struct {
	clientRepository repository.ClientRepository
}

func NewDeleteClient(clientRepository repository.ClientRepository) *DeleteClient {
	return &DeleteClient{
		clientRepository: clientRepository,
	}
}

func (d *DeleteClient) Execute(id string) error {
	existingUser, err := d.clientRepository.FindById(id)

	if err != nil {
		return err
	}

	if existingUser == nil {
		return ErrClientNotFound
	}

	err = d.clientRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
