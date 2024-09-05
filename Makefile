run-api:
	go run cmd/api/main.go

build-api:
	go build -o bin/api cmd/api/main.go
