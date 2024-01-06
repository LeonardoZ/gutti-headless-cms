# Gutti CMS

A simples headless CMS, focused on simple content management.

## Requirements

- Go 1.21.x or compatible
- Docker
- Docker Compose

## Execute (with docker)

Just run:
> `docker-compose up --build`

## Execute (without docker)

Install dependencies
> `go mod download`

Start the application
> `go run main.go`

Start the application in development mode, with auto-reload
> `air`

Run tests
> `go test -v`

## Config

Create a `.env` file based on `.env.example`. Configure the following variables:

- `API_PORT`: TCP port for the application.