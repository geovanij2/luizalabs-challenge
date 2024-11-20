package repository

import (
	"encoding/json"
	"luizalabs-challenge/domain/entity"
	"net/http"
	"strconv"
)

type ProductRepositoryHttp struct {
}

func NewProductRepositoryHttp() *ProductRepositoryHttp {
	return &ProductRepositoryHttp{}
}

func (r *ProductRepositoryHttp) FindById(id string) (*entity.Product, error) {
	url := "http://challenge-api.luizalabs.com/api/product/" + id
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var product entity.Product

	err = json.NewDecoder(resp.Body).Decode(&product)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryHttp) FindAll(offset uint64) ([]*entity.Product, error) {
	url := "http://challenge-api.luizalabs.com/api/product/?page=" + strconv.FormatUint(offset+1, 10)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var products = []*entity.Product{}

	err = json.NewDecoder(resp.Body).Decode(&products)

	if err != nil {
		return nil, err
	}

	return products, nil
}
