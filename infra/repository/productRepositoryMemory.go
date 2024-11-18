package repository

import (
	"errors"
	"luizalabs-chalenge/domain/entity"
)

type ProductRepositoryMemory struct {
	products map[string]*entity.Product
}

func NewProductRepositoryMemory() *ProductRepositoryMemory {
	return &ProductRepositoryMemory{
		products: make(map[string]*entity.Product),
	}
}

func (r *ProductRepositoryMemory) Create(product *entity.Product) (*entity.Product, error) {
	if _, exists := r.products[product.Id]; exists {
		return nil, errors.New("product already exists")
	}
	r.products[product.Id] = product
	return product, nil
}

func (r *ProductRepositoryMemory) FindById(id string) (*entity.Product, error) {
	product, ok := r.products[id]

	if !ok {
		return nil, nil
	}

	return product, nil
}

func (r *ProductRepositoryMemory) FindAll(offset, limit uint64) ([]*entity.Product, error) {
	var products []*entity.Product

	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepositoryMemory) Update(product *entity.Product) (*entity.Product, error) {
	r.products[product.Id] = product

	return product, nil
}

func (r *ProductRepositoryMemory) Delete(id string) error {
	delete(r.products, id)

	return nil
}
