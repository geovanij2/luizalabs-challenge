package application

import "luizalabs-chalenge/domain/repository"

type DeleteFavoriteProduct struct {
	clientRepository    repository.ClientRepository
	productRepository   repository.ProductRepository
	favoritesRepository repository.FavoritesRepository
}

type DeleteFavoriteProductInput struct {
	ClientId  string
	ProductId string
}

func NewDeleteFavoriteProduct(
	clientRepository repository.ClientRepository,
	productRepository repository.ProductRepository,
	favoritesRepository repository.FavoritesRepository,
) *DeleteFavoriteProduct {
	return &DeleteFavoriteProduct{
		clientRepository:    clientRepository,
		productRepository:   productRepository,
		favoritesRepository: favoritesRepository,
	}
}

func (d *DeleteFavoriteProduct) Execute(input DeleteFavoriteProductInput) error {
	client, err := d.clientRepository.FindById(input.ClientId)

	if err != nil {
		return err
	}

	if client == nil {
		return ErrClientNotFound
	}

	product, err := d.productRepository.FindById(input.ProductId)

	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	isFavorite, err := d.favoritesRepository.IsFavorite(input.ClientId, input.ProductId)

	if err != nil {
		return err
	}

	if !isFavorite {
		return ErrProductIsNotFavorite
	}

	err = d.favoritesRepository.RemoveFavorite(input.ClientId, input.ProductId)

	if err != nil {
		return err
	}

	return nil
}
