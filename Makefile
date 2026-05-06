APP_NAME := ses-journal
BIN_DIR := ./bin

.PHONY: run build test clean lint tidy

run:
	go run ./cmd/main.go

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/main.go

test:
	go test -v -race -coverprofile=coverage.out ./...

lint:
	golangci-lint run

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
	rm -f coverage.out
