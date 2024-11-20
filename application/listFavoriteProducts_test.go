package application

import (
	"luizalabs-challenge/domain/entity"
	"luizalabs-challenge/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFavoriteProductsSucess(t *testing.T) {
	favoriteProductsRepository := repository.NewFavoritesRepositoryMemory()
	ListFavoriteProducts := NewListFavoriteProducts(favoriteProductsRepository)

	p1 := entity.Product{
		Id:          "1",
		Brand:       "Brand 1",
		Image:       "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Title:       "Product 1",
		ReviewScore: 4.5,
		Price:       100,
	}

	p2 := entity.Product{
		Id:          "2",
		Brand:       "Brand 2",
		Image:       "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Title:       "Product 2",
		ReviewScore: 3.1,
		Price:       29990,
	}

	favoriteProductsRepository.AddFavorite("1", &p1)
	favoriteProductsRepository.AddFavorite("1", &p2)

	products, err := ListFavoriteProducts.Execute(ListFavoriteProductsInput{
		ClientId: "1",
		Offset:   0,
		Limit:    10,
	})

	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, p1, *products[0])
	assert.Equal(t, p2, *products[1])

	products, err = ListFavoriteProducts.Execute(ListFavoriteProductsInput{
		ClientId: "1",
		Offset:   1,
		Limit:    1,
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, p2, *products[0])
}
