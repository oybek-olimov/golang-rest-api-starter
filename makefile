APP_NAME=go-crud-api
APP_PORT=8080

.PHONY: run build swag docker-up docker-down tidy fmt

run:
	go run cmd/main.go

build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

swag:
	swag init --generalInfo cmd/main.go --output docs

tidy:
	go mod tidy

fmt:
	go fmt ./...

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
