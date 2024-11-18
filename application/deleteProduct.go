package application

import (
	"luizalabs-chalenge/domain/repository"
)

type DeleteProduct struct {
	productRepository repository.ProductRepository
}

func NewDeleteProduct(productRepository repository.ProductRepository) *DeleteProduct {
	return &DeleteProduct{
		productRepository: productRepository,
	}
}

func (d *DeleteProduct) Execute(id string) error {
	product, err := d.productRepository.FindById(id)

	if err != nil {
		return err
	}

	if product == nil {
		return ErrProductNotFound
	}

	err = d.productRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
