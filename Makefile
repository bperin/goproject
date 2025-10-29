.PHONY: help install deps run build test test-verbose clean fmt vet lint all

# Default target - show help
help:
	@echo "Available targets:"
	@echo "  make install       - Install dependencies and tidy go.mod"
	@echo "  make deps          - Alias for install"
	@echo "  make run           - Run the application"
	@echo "  make build         - Build the application binary"
	@echo "  make test          - Run all tests"
	@echo "  make test-verbose  - Run tests with verbose output"
	@echo "  make fmt           - Format code with go fmt"
	@echo "  make vet           - Run go vet for code analysis"
	@echo "  make lint          - Run fmt and vet together"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make all           - Run fmt, vet, test, and build"

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies installed successfully!"

# Alias for install
deps: install

# Run the application
run:
	@echo "Running application..."
	go run main.go

# Build the application
build:
	@echo "Building application..."
	go build -o goproject main.go
	@echo "Build complete! Binary: ./goproject"

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Run tests with verbose output
test-verbose:
	@echo "Running tests (verbose)..."
	go test -v ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Lint (format + vet)
lint: fmt vet
	@echo "Linting complete!"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f goproject
	go clean
	@echo "Clean complete!"

# Run all checks and build
all: lint test build
	@echo "All tasks complete!"
