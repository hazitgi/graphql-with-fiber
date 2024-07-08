# Run application
run:
	go run ./cmd/main.go

build:
	go build -o ./bin/main ./cmd/main.go
# go build -a -installsuffix cgo -o ./bin/main ./cmd/main.go

clean:
	go clean
	rm -f ./bin/main

fmt:
	go fmt ./...

# Install dependencies
deps:
	go get ./...

# Test
test: go test -v ./...


dev:
	air -c .air.toml

up: 
	docker-compose up --build

# Down docker container
down: 
	docker-compose down

