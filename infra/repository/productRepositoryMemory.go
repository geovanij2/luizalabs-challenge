package repository

import (
	"luizalabs-challenge/domain/entity"
)

type ProductRepositoryMemory struct {
	products map[string]*entity.Product
}

func NewProductRepositoryMemory() *ProductRepositoryMemory {
	products := make(map[string]*entity.Product)

	products["1"] = &entity.Product{
		Id:          "1",
		Brand:       "Brand 1",
		Image:       "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Title:       "Product 1",
		ReviewScore: 4.5,
		Price:       100,
	}

	products["2"] = &entity.Product{
		Id:          "2",
		Brand:       "Brand 2",
		Image:       "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Title:       "Product 2",
		ReviewScore: 3.1,
		Price:       29990,
	}

	products["3"] = &entity.Product{
		Id:          "3",
		Brand:       "Brand 2",
		Image:       "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Title:       "Product 3",
		ReviewScore: 3,
		Price:       10100,
	}

	return &ProductRepositoryMemory{
		products: products,
	}
}

func (r *ProductRepositoryMemory) FindById(id string) (*entity.Product, error) {
	product, ok := r.products[id]

	if !ok {
		return nil, nil
	}

	return product, nil
}

func (r *ProductRepositoryMemory) FindAll(offset uint64) ([]*entity.Product, error) {
	var products []*entity.Product

	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}
