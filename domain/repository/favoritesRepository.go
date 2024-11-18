package repository

import "luizalabs-chalenge/domain/entity"

type FavoritesRepository interface {
	AddFavorite(clientId, productId string) error
	RemoveFavorite(clientId, productId string) error
	FindFavoritesByClientId(clientId string, offset, limit uint64) ([]*entity.Product, error)
	IsFavorite(clientId, productId string) (bool, error)
}
