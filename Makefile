run:
	go run main.go

test:
	go test -cover ./...

build:
	go mod download

tidy:
	go mod tidy
