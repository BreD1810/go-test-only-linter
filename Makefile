build:
	go build cmd/go-test-only-linter/go-test-only-linter.go

run:
	@go run cmd/go-test-only-linter/go-test-only-linter.go

install:
	go install cmd/go-test-only-linter/go-test-only-linter.go

.PHONY: build run
