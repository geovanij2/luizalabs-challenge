package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/cryptography"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClientSucess(t *testing.T) {
	repo := repository.NewClientRepositoryMemory()
	bcryptAdapter := cryptography.NewBcryptAdapter()
	createClient := NewCreateClient(repo, bcryptAdapter)
	input := CreateClientInput{
		Name:     "Teste",
		Email:    "foo@bar.com",
		Password: "123456",
	}
	client, err := createClient.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, input.Name, client.Name)
	assert.Equal(t, input.Email, client.Email)
	ok, err := bcryptAdapter.Compare(input.Password, client.Password)
	assert.Equal(t, true, ok)
	assert.Nil(t, err)
}

func TestCreateClientEmailAlreadyExists(t *testing.T) {
	repo := repository.NewClientRepositoryMemory()
	bcryptAdapter := cryptography.NewBcryptAdapter()
	createClient := NewCreateClient(repo, bcryptAdapter)
	input := CreateClientInput{
		Name:     "Teste",
		Email:    "foo@bar.com",
		Password: "123456",
	}
	repo.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})
	client, err := createClient.Execute(input)
	assert.Equal(t, ErrClientEmailAlreadyExists, err)
	assert.Nil(t, client)
}
