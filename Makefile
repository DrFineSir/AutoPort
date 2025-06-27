# Makefile for AutoMap

APP_NAME=AutoMap
GO_FILES=main.go

.PHONY: all build clean release

all: build

build:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)-linux $(GO_FILES)
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o $(APP_NAME)-mac $(GO_FILES)
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME)-windows.exe $(GO_FILES)

clean:
	rm -f $(APP_NAME)-linux $(APP_NAME)-mac $(APP_NAME)-windows.exe

release: build
	@echo "Release build complete."

