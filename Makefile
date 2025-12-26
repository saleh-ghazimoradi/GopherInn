docker-up:
	docker compose up -d

docker-down:
	docker compose down

vet:
	go vet ./...

fmt:
	go fmt ./...

run: fmt vet
	go run . run


