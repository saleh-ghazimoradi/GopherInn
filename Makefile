docker-up:
	docker compose up -d

docker-down:
	docker compose down

vet:
	go vet ./...

fmt:
	go fmt ./..

run: vet fmt
	go run . run

