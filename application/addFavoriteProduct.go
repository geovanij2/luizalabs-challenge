package application

import "luizalabs-chalenge/domain/repository"

type AddFavoriteProduct struct {
	clientRepository    repository.ClientRepository
	productRepository   repository.ProductRepository
	favoritesRepository repository.FavoritesRepository
}

type AddFavoriteProductInput struct {
	clientId  string
	productId string
}

func NewAddFavoriteProduct(clientRepository repository.ClientRepository, productRepository repository.ProductRepository, favoritesRepository repository.FavoritesRepository) *AddFavoriteProduct {
	return &AddFavoriteProduct{
		clientRepository:    clientRepository,
		productRepository:   productRepository,
		favoritesRepository: favoritesRepository,
	}
}

func (a *AddFavoriteProduct) Execute(input AddFavoriteProductInput) error {
	client, err := a.clientRepository.FindById(input.clientId)

	if err != nil {
		return err
	}

	if client == nil {
		return ErrClientNotFound
	}

	product, err := a.productRepository.FindById(input.productId)

	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	isFavorite, err := a.favoritesRepository.IsFavorite(input.clientId, input.productId)

	if err != nil {
		return err
	}

	if isFavorite {
		return ErrProductIsAlreadyClientFavorite
	}

	err = a.favoritesRepository.AddFavorite(input.clientId, input.productId)

	if err != nil {
		return err
	}

	return nil
}
