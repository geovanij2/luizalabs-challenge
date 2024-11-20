package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteClientSucess(t *testing.T) {
	userRepository := repository.NewClientRepositoryMemory()
	deleteClient := NewDeleteClient(userRepository)
	userRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})
	err := deleteClient.Execute("1")
	assert.Nil(t, err)

	client, err := userRepository.FindById("1")
	assert.Nil(t, err)
	assert.Nil(t, client)
}
