package repository

import "luizalabs-challenge/domain/entity"

type FavoritesRepositoryMemory struct {
	favoriteProducts map[string][]*entity.Product
}

func NewFavoritesRepositoryMemory() *FavoritesRepositoryMemory {
	return &FavoritesRepositoryMemory{
		favoriteProducts: make(map[string][]*entity.Product),
	}
}

func (f *FavoritesRepositoryMemory) AddFavorite(clientId string, product *entity.Product) error {
	if _, ok := f.favoriteProducts[clientId]; !ok {
		f.favoriteProducts[clientId] = []*entity.Product{}
	}
	f.favoriteProducts[clientId] = append(f.favoriteProducts[clientId], product)
	return nil
}

func (f *FavoritesRepositoryMemory) RemoveFavorite(clientId, productId string) error {
	if _, ok := f.favoriteProducts[clientId]; !ok {
		return nil
	}

	for i, product := range f.favoriteProducts[clientId] {
		if product.Id == productId {
			f.favoriteProducts[clientId] = append(f.favoriteProducts[clientId][:i], f.favoriteProducts[clientId][i+1:]...)
			return nil
		}
	}
	return nil
}

func (f *FavoritesRepositoryMemory) FindFavoritesByClientId(clientId string, offset, limit uint64) ([]*entity.Product, error) {
	if _, ok := f.favoriteProducts[clientId]; !ok {
		f.favoriteProducts[clientId] = []*entity.Product{}
	}
	return f.favoriteProducts[clientId], nil
}

func (f *FavoritesRepositoryMemory) IsFavorite(clientId, productId string) (bool, error) {
	if _, ok := f.favoriteProducts[clientId]; !ok {
		return false, nil
	}
	for _, product := range f.favoriteProducts[clientId] {
		if product.Id == productId {
			return true, nil
		}
	}
	return false, nil
}
