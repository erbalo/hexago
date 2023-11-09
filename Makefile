GO ?= go
CLI_PARAM ?= http

MAIN = cmd/${CLI_PARAM}/main.go
BIN_DIR = ./bin
CLI_DIR = ${BIN_DIR}/hexago
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

param-selector:
	@if [ -z "$(CLI_PARAM)" ]; then \
		echo "No parameter provided, using default"; \
		echo "Parameter is $(CLI_PARAM)"; \
	else \
		echo "Parameter provided"; \
		echo "Parameter is $(CLI_PARAM)"; \
	fi

build-linux: param-selector
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GO) build -o $(CLI_DIR)-linux -v $(MAIN)

build-darwin: param-selector
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin $(GO) build -o $(CLI_DIR)-darwin -v $(MAIN)

build-windows: param-selector
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows $(GO) build -o $(CLI_DIR)-windows.exe -v $(MAIN)

run: build-linux build-darwin build-windows
	${EXEC}

clean:
	go clean
	rm -rf $(BIN_DIR)

.PHONY: all dependencies run clean
