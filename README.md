# Tasks API

Esta é uma API RESTful construída para solidificar conhecimentos de Go, focando na Arquitetura Hexagonal (Ports and Adapters). A API funciona com um banco de dados PostgreSQL, e implementa um CRUD de 'tasks'.

## Tecnologias utilizadas

- Go 1.23.2
- Docker e Docker Compose
- PostgreSQL
- Echo, pgxpool, godotenv, tern

## Requisitos

- Go 1.23 ou superior
- Docker e Docker Compose
- Tern para migrations
- Banco de dados PostgreSQL (local ou container)


## Variáveis de ambiente

O projeto utilizar um arquivo .env para carregar as variáveis de exemplo. Utilize o arquivo `.env.example` para configurar o ambiente antes de rodar a aplicação.

## Como rodar a aplicação

### Rodando localmente
1. Clone o repositório
2. Instale as dependências do Go: `go mod download`
3. Configure seu arquivo `.env` (crie caso necessário)
4. Com o banco postgres rodando, compile e execute a aplicação: `go run ./cmd/app/main.go`. A aplicação estará rodando na porta indicada no .env (padrão é :8000)
5. Para aplicar as migrations, use o comando `tern migrate --config ./config/migrations/tern.conf --migrations ./config/migrations`

### Rodando com Docker
1. Após clonar o repositório, rode `docker-compose up --build`
2. As migrations devem ser executadas automaticamente no novo banco de dados containerizado. 
3. Depois disso, o servidor estará rodando na porta indicada no `.env` (padrão :8000).

## Endpoints da API
- `POST /tasks`: Cria uma nova tarefa. Fornecer `title` e `description`
- `GET /tasks/:id`: Busca uma tarefa no banco. Fornecer `id` via parametro na url.
