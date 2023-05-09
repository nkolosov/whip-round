include .env
export

docker-build:
	docker build -t $(APP_IMAGE_NAME) .

docker-run:
	docker run --name $(APP_IMAGE_NAME) -p $(APP_PORT):8080 $(APP_IMAGE_NAME)

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-restart:
	docker-compose down
	docker-compose up -d

swag:
	swag init -g cmd/main.go

deps:
	go mod download

build: deps
	go build -o $(APP_IMAGE_NAME) cmd/main.go

run: build
	./$(APP_IMAGE_NAME)

test: deps
	go generate ./...
	go test -v ./...