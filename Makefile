pre-commit:
	go mod tidy

test:
	go test .\... -v

quality:
	go test .\... -cover -v
