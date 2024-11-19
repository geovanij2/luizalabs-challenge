package repository

import (
	"luizalabs-chalenge/domain/entity"
)

type ProductRepository interface {
	FindById(id string) (*entity.Product, error)
	FindAll(offset uint64) ([]*entity.Product, error)
}
