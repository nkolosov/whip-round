include .env
export
# go build
# go mod init
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

test:
	go test -v ./...