package repository

import (
	"luizalabs-chalenge/domain/entity"
)

type ClientRepository interface {
	Create(client *entity.Client) (*entity.Client, error)
	FindByEmail(email string) (*entity.Client, error)
	FindById(clientId string) (*entity.Client, error)
	Update(client *entity.Client) (*entity.Client, error)
	Delete(clientId string) error
}
