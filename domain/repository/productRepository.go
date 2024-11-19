package repository

import (
	"luizalabs-challenge/domain/entity"
)

type ProductRepository interface {
	FindById(id string) (*entity.Product, error)
	FindAll(offset uint64) ([]*entity.Product, error)
}
