# go-expert-exercicios
Projeto destinado a entrega de atividades em Go.

# Como executar o projeto

## Banco de dados
* Abrir um terminal na raiz do projeto
* cd docker-db
* docker-compose up -d

Vai rodar uma instância de MariaDB com 2 pedidos cadastrados.

## coloca a aplicação no ar

* go run main.go

O comando vai colocar a aplicação no ar nas seguintes portas:
* REST: porta 8081
* gRPC: porta 8082
* GraphQL: porta 8083

# Como validar
## REST
    * Disponível no arquivo que está na raiz do projeto:
        * api.http

## gRPC
* cd grpc-config/client
* go run client.go

## GraphQL
* Disponível no arquivo que está na raiz do projeto:
    * api.http


<br/>
<br/>
<br/>
<br/>
<br/>

# Bibliotecas Utilizadas


# Requisitos

É necessário criar um usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

# Bibliotecas Utilizadas

## Banco de dados MariaDB
* go get -u github.com/go-sql-driver/mysql

## validator - para validar as structs
* go get github.com/go-playground/validator/v10

## GRPC
* go get google.golang.org/grpc

## GraphQL
* go get github.com/graphql-go/graphql