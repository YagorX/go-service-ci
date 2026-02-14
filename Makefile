.PHONY: lint test build docker-build docker-push docker-run test-integration e2e-up-local e2e-down-local e2e-up-ci e2e-down-ci test-e2e-docker

GITHUB_USER ?= yagorx
IMAGE_NAME ?= ghcr.io/$(shell echo $(GITHUB_USER) | tr '[:upper:]' '[:lower:]')/go-service-ci
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

e2e-up-local:
	bash ./scripts/up-local.sh

e2e-down-local:
	docker compose --env-file ./.env -f e2e_tests/docker-compose.e2e.yml --profile local down -v

e2e-up-ci:
	docker compose --env-file ./.env -f e2e_tests/docker-compose.e2e.yml --profile ci up -d --wait app

e2e-down-ci:
	docker compose --env-file ./.env -f e2e_tests/docker-compose.e2e.yml --profile ci down -v

test-e2e-docker:
	bash ./scripts/up-e2e.sh
