# Makefile for SidVM project

# Name of the Go file and output binary
GO_FILE = sidvm.go
OUTPUT = sidvm

# Go commands
GO = go

# Default target to build the VM
all: build

# Build the Go program
build:
	$(GO) build -o $(OUTPUT) $(GO_FILE)

# Run the Go program
run: build
	./$(OUTPUT)

# Clean up compiled files
clean:
	rm -f $(OUTPUT)

# Install dependencies (if needed)
install:
	$(GO) install

# Format Go files
fmt:
	$(GO) fmt

# Lint the Go code (requires golint)
lint:
	golint $(GO_FILE)

.PHONY: all build run clean install fmt lint