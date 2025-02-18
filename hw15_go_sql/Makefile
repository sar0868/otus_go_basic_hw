# https://github.com/swaggo/gin-swagger

lint: 
	golangci-lint run ./...

test:
	go test ./...

run:
	go run cmd/main.go

swag-init:
	swag init -g cmd/main.go
	