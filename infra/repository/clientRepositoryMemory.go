package repository

import (
	"errors"
	"luizalabs-chalenge/domain/entity"
)

type ClientRepositoryMemory struct {
	clientsById    map[string]*entity.Client
	clientsByEmail map[string]*entity.Client
}

func (c *ClientRepositoryMemory) FindById(clientId string) (*entity.Client, error) {
	client, exists := c.clientsById[clientId]
	if !exists {
		return nil, nil
	}
	return client, nil
}

func NewClientRepositoryMemory() *ClientRepositoryMemory {
	return &ClientRepositoryMemory{
		clientsById:    make(map[string]*entity.Client),
		clientsByEmail: make(map[string]*entity.Client),
	}
}

func (c *ClientRepositoryMemory) Delete(clientId string) error {
	client, exists := c.clientsById[clientId]
	if !exists {
		return nil
	}
	delete(c.clientsById, clientId)
	delete(c.clientsByEmail, client.Email)
	return nil
}

func (c *ClientRepositoryMemory) FindByEmail(email string) (*entity.Client, error) {
	client, exists := c.clientsByEmail[email]
	if !exists {
		return nil, nil
	}
	return client, nil
}

func (c *ClientRepositoryMemory) Update(client *entity.Client) (*entity.Client, error) {
	c.clientsById[client.Id] = client
	c.clientsByEmail[client.Email] = client
	return client, nil
}

func (c *ClientRepositoryMemory) Create(client *entity.Client) (*entity.Client, error) {
	_, exists := c.clientsById[client.Id]
	if exists {
		return nil, errors.New("client ja existe")
	}
	c.clientsById[client.Id] = client
	c.clientsByEmail[client.Email] = client
	return client, nil
}
