include .env
export

### Docker ###
.PHONY: docker-up docker-down
docker-up:
	docker-compose up -d
docker-down:
	docker-compose down

### Dependencies ###
.PHONY: deps swag migrate-install
deps:
	go mod download
migrate-install:
	curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
swag:
	swag init -g cmd/main.go

### Build and Run ###
.PHONY: build run
build: deps
	go build -o $(APP_IMAGE_NAME) cmd/main.go
run: build
	./$(APP_IMAGE_NAME)

### Tests ###
.PHONY: test
test: deps
	go generate ./...
	go test -v ./...

### Migrations ###
.PHONY: migrate-up migrate-down
migrate-up:
	migrate -path ./schema -database 'postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)' -verbose up
migrate-down:
	migrate -path ./schema -database 'postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)' -verbose down

### Linter ###
.PHONY: .install-linter lint lint-fast
GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint

.install-linter:
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.46.2
lint: .install-linter
	$(GOLANGCI_LINT) run ./... --config=./.golangci.yml
lint-fast: .install-linter
	$(GOLANGCI_LINT) run ./... --fast --config=./.golangci.yml