# API Project

Este repositório contém duas APIs: `user-api` e `order-api`. Ambas as APIs são construídas com Golang e utilizam o framework Gin. As APIs estão documentadas com Swagger e são executadas em containers Docker.

## Sumário

- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Execução](#execução)
- [APIs](#apis)
  - [User API](#user-api)
  - [Order API](#order-api)
- [Documentação via Swagger](#documentação-via-swagger)
- [Estrutura do Projeto](#estrutura-do-projeto)

## Pré-requisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.21.5+](https://golang.org/)

## Instalação

Clone o repositório:

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio

## Execução
Para compilar e executar os containers Docker, use o comando em cada pasta:

```bash
Copy code
docker network inspect user-order-network
docker-compose up --build



# APIs
## USER API

A user-api fornece endpoints para gerenciar usuários. A API é exposta na porta 8081.

ENDPOINTS
GET /users: Retorna todos os usuários
GET /users/:id: Retorna um usuário específico pelo ID
POST /users: Cria um novo usuário
PUT /users/:id: Atualiza um usuário existente pelo ID
DELETE /users/:id: Deleta um usuário pelo ID

## ORDER API

A order-api fornece endpoints para gerenciar pedidos. A API é exposta na porta 8080.

ENDPOINTS
GET /orders: Retorna todos os pedidos
GET /orders/:id: Retorna um pedido específico pelo ID
GET /users/:id/orders: Retorna todos os pedidos de um usuário específico
POST /orders: Cria um novo pedido
PUT /orders/:id: Atualiza um pedido existente pelo ID
DELETE /orders/:id: Deleta um pedido pelo ID

## Documentação via Swagger
A documentação do Swagger para ambas as APIs está disponível nos URLs abaixo:

User API: http://localhost:8081/swagger/index.html
Order API: http://localhost:8080/swagger/index.html

## Estrutura do Projeto
O repositório está organizado da seguinte forma:

.
├── user-api
│   ├── controllers
│   │   └── userController.go
│   ├── models
│   │   └── user.go
│   ├── routes
│   │   └── userRoutes.go
│   ├── services
│   │   └── userService.go
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── docs
├── order-api
│   ├── controllers
│   │   └── orderController.go
│   ├── models
│   │   └── order.go
│   ├── routes
│   │   └── orderRoutes.go
│   ├── services
│   │   └── orderService.go
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── docs
├── docker-compose.yml
└── README.md
