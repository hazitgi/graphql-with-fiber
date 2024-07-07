# Run application
run:
	go run ./cmd/main.go

build:
	go build -a -installsuffix cgo -o ./bin/main ./cmd/main.go

fmt:
	go fmt ./...

dev:
	air