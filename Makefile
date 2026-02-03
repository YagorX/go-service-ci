.PHONY: lint test build docker-build docker-push docker-run test-integration

GITHUB_USER ?= yagorx
IMAGE_NAME ?= ghcr.io/$(shell echo $(GITHUB_USER) | tr '[:upper:]' '[:lower:]')/go-ci-cd-demo
IMAGE_TAG ?= local
PLATFORM ?= linux/amd64

APP_NAME := app
BIN_DIR := bin

lint:
	golangci-lint run -v

unit-test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

test-integration:
	docker compose -f ./test/docker-compose-test.yml up -d --wait
	go test --tags=test_integration ./test/...
	docker compose -f ./test/docker-compose-test.yml down

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd

docker-build:
	docker buildx build \
		--platform $(PLATFORM) \
		-t $(IMAGE_NAME):$(IMAGE_TAG) \
		.

docker-push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)

docker-run:
	docker run --rm -p 8080:8080 $(IMAGE_NAME):$(IMAGE_TAG)

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

