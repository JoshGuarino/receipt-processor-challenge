run:
	go run main.go

test:
	go test -cover ./internal/handlers/... ./internal/services/...

build:
	go mod download

tidy:
	go mod tidy
