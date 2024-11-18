package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"
)

type ListProducts struct {
	productRepository repository.ProductRepository
}

type ListProductsInput struct {
	Offset uint64
	Limit  uint64
}

func NewListProducts(productRepository repository.ProductRepository) *ListProducts {
	return &ListProducts{
		productRepository: productRepository,
	}
}

func (l *ListProducts) Execute(input ListFavoriteProductsInput) ([]*entity.Product, error) {
	products, err := l.productRepository.FindAll(input.Offset, input.Limit)

	if err != nil {
		return nil, err
	}

	return products, nil
}
