docker-up:
	@docker compose up -d

docker-down:
	@docker compose down

vet:
	@go vet ./...

fmt:
	@go fmt ./...

test:
	@go test ./...

build:
	mkdir -p bin
	@go build -o bin/gopherinn

run: fmt vet
	@go run . run


