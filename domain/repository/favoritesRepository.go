package repository

import "luizalabs-challenge/domain/entity"

type FavoritesRepository interface {
	AddFavorite(clientId string, product *entity.Product) error
	RemoveFavorite(clientId, productId string) error
	FindFavoritesByClientId(clientId string, offset, limit uint64) ([]*entity.Product, error)
	IsFavorite(clientId, productId string) (bool, error)
}
