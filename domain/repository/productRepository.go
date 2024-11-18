package repository

import (
	"luizalabs-chalenge/domain/entity"
)

type ProductRepository interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	FindAll(offset, limit uint64) ([]*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id string) error
}
