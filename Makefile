
.PHONY: all
all: test

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: install-deps
install-deps: ## golangci-lint and go mod dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	go mod tidy
	go mod download
