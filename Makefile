run:
	go run ./cmd/main.go

build:
	go build -o main ./cmd/main.go

DEFAULT_GOAL: run