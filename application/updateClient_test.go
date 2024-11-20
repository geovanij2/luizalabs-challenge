package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateClientSuccess(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	updateClient := NewUpdateClient(clientRepository)

	client := entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	}
	clientRepository.Create(&client)
	newName := "Bruno"
	client.Name = newName
	err := updateClient.Execute(&client)
	assert.Nil(t, err)

	updatedClient, err := clientRepository.FindById(client.Id)
	assert.Nil(t, err)
	assert.Equal(t, client.Id, updatedClient.Id)
	assert.Equal(t, client.Email, updatedClient.Email)
	assert.Equal(t, newName, updatedClient.Name)
	assert.Equal(t, client.Password, updatedClient.Password)
}

func TestUpdateClientEmailAlreadyExists(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	updateClient := NewUpdateClient(clientRepository)

	client := entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	}

	anotherClient := entity.Client{
		Id:       "2",
		Name:     "Bruno",
		Email:    "foo@bar2.com",
		Password: "123456",
	}

	clientRepository.Create(&client)
	clientRepository.Create(&anotherClient)

	input := entity.Client{
		Id:    "1",
		Email: "foo@bar2.com",
	}
	err := updateClient.Execute(&input)
	assert.Equal(t, ErrClientEmailAlreadyExists, err)
}
