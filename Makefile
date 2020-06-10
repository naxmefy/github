.PHONY: run
run:
	@go run cmd/github/main.go

.PHONY: test
test:
	@go test -cover ./...

.PHONY: install
install:
	@go install ./...

.PHONY: build
build:
	@go build ./...
