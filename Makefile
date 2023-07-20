build:
	go mod tidy
	go vet ./...
	go build -v
