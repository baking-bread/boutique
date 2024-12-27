pre-commit:
	go mod tidy
	go vet ./...

test:
	go test ./... -v

quality:
	go test ./... -cover -v
