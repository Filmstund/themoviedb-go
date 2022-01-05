PROJECT_NAME := themoviedb-go
GOFMT_FILES = $(shell go list -f '{{.Dir}}' ./... | grep -v '/pb')

all: \
	mod-tidy \
	fmt \
	lint \
	test \
	build \
	verify-nodiff
.PHONY: all

lint:
	$(info [$@] linting $(PROJECT_NAME)...)
	@golangci-lint run --config .golangci.yaml
.PHONY: lint

verify-nodiff:
	$(info [$@] verifying no git diff...)
	@git update-index --refresh && git diff-index --quiet HEAD --
.PHONY: verify-nodiff

# Format all files
fmt:
	$(info [$@] formatting all Go files...)
	@gofumpt -w $(GOFMT_FILES)
.PHONY: fmt

mod-tidy:
	$(info [$@] tidying up...)
	@go mod tidy
.PHONY: mod-tidy

test:
	$(info [$@] running the tests...)
	@go test \
		-shuffle=on \
		-count=1 \
		-short \
		-timeout=5m \
		./...
.PHONY: test

build:
	$(info [$@] building ${PROJECT_NAME}...)
	@go build ./...
.PHONY: build
