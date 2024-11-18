package application

import (
	"luizalabs-chalenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClientSucess(t *testing.T) {
	repo := repository.NewClientRepositoryMemory()
	createClient := NewCreateClient(repo)
	input := CreateClientInput{
		Name:     "Teste",
		Email:    "foo@bar.com",
		Password: "123456",
	}
	client, err := createClient.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, input.Name, client.Name)
	assert.Equal(t, input.Email, client.Email)
	assert.Equal(t, input.Password, client.Password)
}
