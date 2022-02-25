make:
	go run main.go

test:
	go test ./... -short

build:
	go build -o bubbletea-starter main.go

install:
	go install