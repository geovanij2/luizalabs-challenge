package repository

import (
	"luizalabs-challenge/domain/entity"

	"github.com/jmoiron/sqlx"
)

type ClientRepositoryDatabase struct {
	db *sqlx.DB
}

func NewClientRepositoryDatabase(db *sqlx.DB) *ClientRepositoryDatabase {
	return &ClientRepositoryDatabase{
		db: db,
	}
}

func (c *ClientRepositoryDatabase) Delete(clientId string) error {
	_, err := c.db.Exec("DELETE FROM clients WHERE id = $1", clientId)

	if err != nil {
		return err
	}

	return nil
}

func (c *ClientRepositoryDatabase) FindById(clientId string) (*entity.Client, error) {
	var client entity.Client
	c.db.Get(&client, "SELECT * FROM clients WHERE id = $1", clientId)

	if client.Id == "" {
		return nil, nil
	}

	return &client, nil
}

func (c *ClientRepositoryDatabase) FindByEmail(email string) (*entity.Client, error) {
	var client entity.Client
	c.db.Get(&client, "SELECT * FROM clients WHERE email = $1", email)

	if client.Id == "" {
		return nil, nil
	}

	return &client, nil
}

func (c *ClientRepositoryDatabase) Update(client *entity.Client) (*entity.Client, error) {
	_, err := c.db.Exec("UPDATE clients SET name = $1, email = $2, password = $3 WHERE id = $4", client.Name, client.Email, client.Password, client.Id)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientRepositoryDatabase) Create(client *entity.Client) (*entity.Client, error) {
	_, err := c.db.Exec("INSERT INTO clients (id, name, email, password) VALUES ($1, $2, $3, $4)", client.Id, client.Name, client.Email, client.Password)
	if err != nil {
		return nil, err
	}
	return client, nil
}
