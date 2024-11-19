# Luizalabs-challenge

## Como rodar

### Pré-requisitos

- Docker
- Docker Compose

### Executando

```bash
docker-compose up -d
```

Esse comando irá iniciar o projeto em um container docker e iniciar o banco de dados e o servidor de aplicação.
Ele também irá iniciar o PostgreSQL e o pgAdmin para que você possa interagir com o banco de dados.

### Como utilizar a API

Para interagir com a API, você pode usar o Postman ou qualquer outro cliente HTTP que você preferir.

Estou disponibilizando uma collection do postman para você poder testar a API.

basta baixar o arquivo *luizalabs-challenge.postman_collection.json* e importar no postman.

### Rotas disponíveis

#### Clientes

- POST /clients
  - Cria um novo cliente
  - Status: 201
  - Body:
    ```json
    {
      "id": "string",
      "name": "string",
      "email": "string",
      "password": "string"
    }
    ```

- GET /clients/{clientId}
  - Retorna o cliente com o id especificado
  - Status: 200
  - Headers:
    - Authorization: Bearer {accessToken}
  - Body:
    ```json
    {
      "id": "string",
      "name": "string",
      "email": "string",
      "password": "string"
    }
    ```
- PUT /clients/{clientId}
  - Atualiza o cliente com o id especificado
  - Pode ser atualizado o nome e/ou email
  - Status: 200
  - Headers:
    - Authorization: Bearer {accessToken}
  - Body:
    ```json
    {
      "name": "string",
      "email": "string",
    }
    ```
- DELETE /clients/{clientId}
  - Deleta o cliente com o id especificado
  - Status: 200
  - Headers:
    - Authorization: Bearer {accessToken}

#### Produtos favoritos

- POST /clients/{clientId}/favorites
  - Adiciona um novo produto aos favoritos do cliente
  - Status: 201
  - Headers:
    - Authorization: Bearer {accessToken}
  - Body:
    ```json
    {
      "productId": "string"
    }
    ```
- GET /clients/{clientId}/favorites
  - Retorna os produtos favoritos do cliente
  - Status: 200
  - Headers:
    - Authorization: Bearer {accessToken}
  - Body:
    ```json
    [
      {
        "id": "string",
        "price": 0,
        "image": "string",
        "brand": "string",
        "title": "string",
        "reviewScore": 0
      }
    ]
    ```
- DELETE /clients/{clientId}/favorites/{productId}
  - Deleta o produto do cliente
  - Status: 200
  - Headers:
    - Authorization: Bearer {accessToken}

#### Login

- POST /login
  - Retorna o accessToken para o cliente
  - Status: 201
  - Body:
    ```json
    {
      "accessToken": "string"
    }
    ```

## Acessando o banco de dados

Para acessar o banco de dados, você pode usar o pgAdmin ou qualquer outro cliente que você preferir.

Basta acessar o endereço http://localhost:5050 e fazer login com o usuário *admin@example.com* e a senha *admin*.

Uma vez logado você pode criar uma nova conexão com o banco de dados. Lembrando que o hostname é *psql*, o username é *postgres* e o password é *123*.