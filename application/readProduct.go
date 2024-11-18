package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"
)

type ReadProduct struct {
	productRepository repository.ProductRepository
}

type ReadProductInput struct {
	ProductId string
}

func NewReadProduct(productRepository repository.ProductRepository) *ReadProduct {
	return &ReadProduct{
		productRepository: productRepository,
	}
}

func (r *ReadProduct) Execute(input ReadProductInput) (*entity.Product, error) {
	product, err := r.productRepository.FindById(input.ProductId)

	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}
