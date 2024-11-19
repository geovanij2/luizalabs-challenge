package repository

import (
	"luizalabs-challenge/domain/entity"

	"github.com/jmoiron/sqlx"
)

type FavoritesRepositoryDatabase struct {
	db *sqlx.DB
}

func NewFavoritesRepositoryDatabase(db *sqlx.DB) *FavoritesRepositoryDatabase {
	return &FavoritesRepositoryDatabase{
		db: db,
	}
}

func (f *FavoritesRepositoryDatabase) AddFavorite(clientId string, product *entity.Product) error {
	_, err := f.db.Exec("INSERT INTO favorite_products (client_id, product_id, price, image, brand, title, review_score) VALUES ($1, $2, $3, $4, $5, $6, $7)", clientId, product.Id, product.Price, product.Image, product.Brand, product.Title, product.ReviewScore)
	if err != nil {
		return err
	}
	return nil
}

func (f *FavoritesRepositoryDatabase) RemoveFavorite(clientId, productId string) error {
	_, err := f.db.Exec("DELETE FROM favorite_products WHERE client_id = $1 AND product_id = $2", clientId, productId)
	if err != nil {
		return err
	}
	return nil
}

func (f *FavoritesRepositoryDatabase) FindFavoritesByClientId(clientId string, offset, limit uint64) ([]*entity.Product, error) {
	products := []*entity.Product{}
	err := f.db.Select(&products, "SELECT product_id, price, image, brand, title, review_score FROM favorite_products WHERE client_id = $1 ORDER BY product_id LIMIT $2 OFFSET $3", clientId, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (f *FavoritesRepositoryDatabase) IsFavorite(clientId, productId string) (bool, error) {
	var product entity.Product
	f.db.Get(&product, "SELECT product_id, price, image, brand, title, review_score FROM favorite_products WHERE client_id = $1 AND product_id = $2", clientId, productId)
	if product.Id == "" {
		return false, nil
	}
	return true, nil
}
