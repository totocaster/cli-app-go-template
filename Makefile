BINARY ?= gocli
CMD_PATH ?= ./cmd/$(BINARY)
BUILD_DIR ?= bin
INSTALL_DIR ?= $(HOME)/.local/bin
GO ?= go
GOLANGCI_LINT ?= golangci-lint

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

.PHONY: build install test test-race fmt lint tidy clean release-snapshot help

build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY) $(CMD_PATH)

install: build
	@mkdir -p $(INSTALL_DIR)
	cp $(BUILD_DIR)/$(BINARY) $(INSTALL_DIR)/$(BINARY)

test:
	$(GO) test ./...

test-race:
	$(GO) test -race ./...

fmt:
	$(GO) fmt ./...

lint:
	$(GOLANGCI_LINT) run

tidy:
	$(GO) mod tidy

clean:
	rm -rf $(BUILD_DIR) dist coverage.out coverage.html

release-snapshot:
	goreleaser release --snapshot --clean

help:
	@echo "make build            Build the CLI into ./bin"
	@echo "make install          Install the CLI into ~/.local/bin"
	@echo "make test             Run unit tests"
	@echo "make test-race        Run tests with the race detector"
	@echo "make fmt              Format the codebase"
	@echo "make lint             Run golangci-lint"
	@echo "make tidy             Tidy go.mod/go.sum"
	@echo "make clean            Remove build artifacts"
	@echo "make release-snapshot Build release artifacts locally with GoReleaser"
