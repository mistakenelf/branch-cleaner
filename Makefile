make:
	go run main.go

test:
	go test ./... -short

build:
	go build -o branch-cleaner main.go

install:
	go install