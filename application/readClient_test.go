package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadClientSucess(t *testing.T) {
	repo := repository.NewClientRepositoryMemory()
	readClient := NewReadClient(repo)

	repo.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	input := ReadClientInput{
		ClientId: "1",
	}
	client, err := readClient.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, "1", client.Id)
	assert.Equal(t, "Joao", client.Name)
	assert.Equal(t, "foo@bar.com", client.Email)
	assert.Equal(t, "123456", client.Password)
}

func TestReadClientNotFound(t *testing.T) {
	repo := repository.NewClientRepositoryMemory()
	readClient := NewReadClient(repo)
	input := ReadClientInput{
		ClientId: "2",
	}
	_, err := readClient.Execute(input)
	assert.Equal(t, ErrClientNotFound, err)
}
