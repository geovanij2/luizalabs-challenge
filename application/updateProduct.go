package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"
)

type UpdateProduct struct {
	productRepository repository.ProductRepository
}

func NewUpdateProduct(productRepository repository.ProductRepository) *UpdateProduct {
	return &UpdateProduct{
		productRepository: productRepository,
	}
}

func (u *UpdateProduct) Execute(prod *entity.Product) error {
	existingProduct, err := u.productRepository.FindById(prod.Id)

	if err != nil {
		return err
	}

	if existingProduct == nil {
		return ErrProductNotFound
	}

	_, err = u.productRepository.Update(prod)

	if err != nil {
		return err
	}

	return nil
}
