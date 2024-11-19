package application

import "luizalabs-challenge/domain/repository"

type AddFavoriteProduct struct {
	clientRepository    repository.ClientRepository
	productRepository   repository.ProductRepository
	favoritesRepository repository.FavoritesRepository
}

type AddFavoriteProductInput struct {
	ClientId  string `json:"clientId"`
	ProductId string `json:"productId"`
}

func NewAddFavoriteProduct(clientRepository repository.ClientRepository, productRepository repository.ProductRepository, favoritesRepository repository.FavoritesRepository) *AddFavoriteProduct {
	return &AddFavoriteProduct{
		clientRepository:    clientRepository,
		productRepository:   productRepository,
		favoritesRepository: favoritesRepository,
	}
}

func (a *AddFavoriteProduct) Execute(input AddFavoriteProductInput) error {
	client, err := a.clientRepository.FindById(input.ClientId)

	if err != nil {
		return err
	}

	if client == nil {
		return ErrClientNotFound
	}

	product, err := a.productRepository.FindById(input.ProductId)

	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	isFavorite, err := a.favoritesRepository.IsFavorite(input.ClientId, input.ProductId)

	if err != nil {
		return err
	}

	if isFavorite {
		return ErrProductIsAlreadyClientFavorite
	}

	err = a.favoritesRepository.AddFavorite(input.ClientId, product)

	if err != nil {
		return err
	}

	return nil
}
