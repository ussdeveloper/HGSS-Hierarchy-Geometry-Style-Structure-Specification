@echo off
REM Build script for Windows

echo Building HGSS Converter for multiple platforms...

REM Create bin directory
if not exist bin mkdir bin

REM Build for Windows
echo Building for Windows AMD64...
go build -o bin/hgss-converter-windows-amd64.exe main.go

echo Building for Windows 386...
go build -o bin/hgss-converter-windows-386.exe main.go

REM Build for Linux
echo Building for Linux AMD64...
set GOOS=linux
set GOARCH=amd64
go build -o bin/hgss-converter-linux-amd64 main.go

echo Building for Linux ARM64...
set GOARCH=arm64
go build -o bin/hgss-converter-linux-arm64 main.go

REM Build for macOS
echo Building for macOS AMD64...
set GOOS=darwin
set GOARCH=amd64
go build -o bin/hgss-converter-mac-amd64 main.go

echo Building for macOS ARM64...
set GOARCH=arm64
go build -o bin/hgss-converter-mac-arm64 main.go

REM Reset environment
set GOOS=
set GOARCH=

echo Build completed! Binaries are in bin/ directory.