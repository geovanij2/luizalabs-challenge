package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/domain/repository"
)

type ListFavoriteProducts struct {
	favoritesRepository repository.FavoritesRepository
}

type ListFavoriteProductsInput struct {
	ClientId string
	Offset   uint64
	Limit    uint64
}

func NewListFavoriteProducts(favoritesRepository repository.FavoritesRepository) *ListFavoriteProducts {
	return &ListFavoriteProducts{
		favoritesRepository: favoritesRepository,
	}
}

func (l *ListFavoriteProducts) Execute(input ListFavoriteProductsInput) ([]*entity.Product, error) {
	products, err := l.favoritesRepository.FindFavoritesByClientId(input.ClientId, input.Offset, input.Limit)

	if err != nil {
		return nil, err
	}

	return products, nil
}
