GO ?= go

MAIN = cmd/main.go
CLI_DIR = ./bin/hexago
EXEC = ${CLI_DIR}

# Detect the operating system
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
  EXEC=${CLI_DIR}-linux
endif

ifeq ($(UNAME_S),Darwin)
  EXEC=${CLI_DIR}-darwin
endif

ifeq ($(OS),Windows_NT)
	EXEC=${CLI_DIR}-windows.exe
else
  ifeq ($(UNAME_S),CYGWIN*)
		EXEC=${CLI_DIR}-windows.exe
  endif
  ifeq ($(UNAME_S),MINGW*)
		EXEC=${CLI_DIR}-windows.exe
  endif
endif


all:
	$(MAKE) dependencies
	$(MAKE) run

dependencies:
	$(GO) mod download
	$(GO) mod tidy

build-linux:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO) build -o $(CLI_DIR)-linux -v $(MAIN)

build-darwin:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GO) build -o $(CLI_DIR)-darwin -v $(MAIN)

build-windows:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows $(GO) build -o $(CLI_DIR)-windows.exe -v $(MAIN)

run: build-linux build-darwin build-windows
	${EXEC}

clean:
	go clean
	rm $(CLI_DIR)-darwin
	rm $(CLI_DIR)-linux
	rm $(CLI_DIR)-windows.exe

.PHONY: all dependencies run clean
