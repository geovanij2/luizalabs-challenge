package application

import "errors"

var (
	ErrClientNotFound                 = errors.New("client not found")
	ErrClientEmailAlreadyExists       = errors.New("email already exists")
	ErrProductAlreadyExists           = errors.New("product already exists")
	ErrProductNotFound                = errors.New("product not found")
	ErrProductIsAlreadyClientFavorite = errors.New("product is already client favorite")
	ErrProductIsNotFavorite           = errors.New("product is not favorite")
	ErrWrongPassword                  = errors.New("wrong password")
)
