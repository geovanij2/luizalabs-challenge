package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteFavoriteProductSucess(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	deleteFavoriteProduct := NewDeleteFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	clientRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	favoritesRepository.AddFavorite("1", &entity.Product{
		Id:          "1",
		Brand:       "Brand 1",
		Image:       "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Title:       "Product 1",
		ReviewScore: 4.5,
		Price:       100,
	})

	isFavorite, err := favoritesRepository.IsFavorite("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, true, isFavorite)

	deleteFavoriteProduct.Execute(DeleteFavoriteProductInput{
		ClientId:  "1",
		ProductId: "1",
	})

	isFavorite, err = favoritesRepository.IsFavorite("1", "1")
	assert.Nil(t, err)
	assert.Equal(t, false, isFavorite)
}

func TestDeleteFavoriteProductNotFound(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	deleteFavoriteProduct := NewDeleteFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	clientRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	err := deleteFavoriteProduct.Execute(DeleteFavoriteProductInput{
		ClientId:  "1",
		ProductId: "NotFound",
	})

	assert.Equal(t, ErrProductNotFound, err)
}

func TestDeleteFavoriteProductClientNotFound(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	deleteFavoriteProduct := NewDeleteFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	err := deleteFavoriteProduct.Execute(DeleteFavoriteProductInput{
		ClientId:  "NotFound",
		ProductId: "1",
	})

	assert.Equal(t, ErrClientNotFound, err)
}

func TestDeleteFavoriteProductProductNotFavorite(t *testing.T) {
	clientRepository := repository.NewClientRepositoryMemory()
	productRepository := repository.NewProductRepositoryMemory()
	favoritesRepository := repository.NewFavoritesRepositoryMemory()
	deleteFavoriteProduct := NewDeleteFavoriteProduct(clientRepository, productRepository, favoritesRepository)

	clientRepository.Create(&entity.Client{
		Id:       "1",
		Name:     "Joao",
		Email:    "foo@bar.com",
		Password: "123456",
	})

	err := deleteFavoriteProduct.Execute(DeleteFavoriteProductInput{
		ClientId:  "1",
		ProductId: "1",
	})

	assert.Equal(t, ErrProductIsNotFavorite, err)
}
