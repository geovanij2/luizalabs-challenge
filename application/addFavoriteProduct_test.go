package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFavoriteProductSucess(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	AddFavoriteProduct := NewAddFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	clientRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	input := AddFavoriteProductInput{
		ClientId:  "1",
		ProductId: "1",
	}

	AddFavoriteProduct.Execute(input)
	isFavorite, err := favoritesRepository.IsFavorite("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, true, isFavorite)
}

func TestAddFavoriteProductNotFound(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	AddFavoriteProduct := NewAddFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	clientRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	input := AddFavoriteProductInput{
		ClientId:  "1",
		ProductId: "NotFound",
	}

	err := AddFavoriteProduct.Execute(input)
	assert.Equal(t, ErrProductNotFound, err)
}

func TestAddFavoriteProductClientNotFound(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	AddFavoriteProduct := NewAddFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	input := AddFavoriteProductInput{
		ClientId:  "NotFound",
		ProductId: "1",
	}

	err := AddFavoriteProduct.Execute(input)
	assert.Equal(t, ErrClientNotFound, err)
}

func TestAddFavoriteProductProductAlreadyClientFavorite(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	AddFavoriteProduct := NewAddFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	clientRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	input := AddFavoriteProductInput{
		ClientId:  "1",
		ProductId: "1",
	}

	AddFavoriteProduct.Execute(input)
	isFavorite, err := favoritesRepository.IsFavorite("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, true, isFavorite)

	err = AddFavoriteProduct.Execute(input)
	assert.Equal(t, ErrProductIsAlreadyClientFavorite, err)
}
