# Makefile for HGSS Converter

.PHONY: all test build clean build-linux build-mac build-windows

# Default target
all: test build

# Run tests
test:
	go test ./...

# Build for current platform
build:
	go build -o bin/hgss-converter main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/hgss-converter-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/hgss-converter-linux-arm64 main.go

# Build for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/hgss-converter-mac-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/hgss-converter-mac-arm64 main.go

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/hgss-converter-windows-amd64.exe main.go
	GOOS=windows GOARCH=386 go build -o bin/hgss-converter-windows-386.exe main.go

# Build for all platforms
build-all: build-linux build-mac build-windows

# Create bin directory
bin:
	mkdir -p bin

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run integration tests
test-integration: build
	./bin/hgss-converter -i testdata/sample.geojson -o test_output.hgss.json -f geojson -t hgss
	./bin/hgss-converter -i testdata/sample.hgss.json -o test_output.geojson -f hgss -t geojson
	./bin/hgss-converter -i testdata/sample.kml -o test_output_kml.hgss.json -f kml -t hgss
	@echo "Integration tests passed"

# Clean test outputs
clean-test:
	rm -f test_output.*