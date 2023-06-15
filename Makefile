build:
	go mod tidy
	go vet ./...
	go build -v github.com/sic-xe/sicer
