package application

import (
	"luizalabs-chalenge/domain/entity"
	"luizalabs-chalenge/domain/repository"

	"github.com/google/uuid"
)

type CreateProduct struct {
	productRepository repository.ProductRepository
}

type CreateProductInput struct {
	Price       uint64
	Image       string
	Brand       string
	Title       string
	ReviewScore float64
}

func NewCreateProduct(productRepository repository.ProductRepository) *CreateProduct {
	return &CreateProduct{
		productRepository: productRepository,
	}
}

func (c *CreateProduct) Execute(input CreateProductInput) (*entity.Product, error) {
	newProduct := entity.Product{
		Id:          uuid.NewString(),
		Price:       input.Price,
		Image:       input.Image,
		Brand:       input.Brand,
		Title:       input.Title,
		ReviewScore: input.ReviewScore,
	}

	product, err := c.productRepository.Create(&newProduct)

	if err != nil {
		return nil, err
	}

	return product, nil
}
