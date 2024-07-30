BINARY_NAME=TaskList
GO := go
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S), Linux)
    OS := linux
    EXT :=
    BIN_DIR := bin/linux
endif
ifeq ($(UNAME_S), Darwin)
    OS := darwin
    EXT :=
    BIN_DIR := bin/darwin
endif
ifeq ($(OS), Windows_NT)
    OS := windows
    EXT := .exe
    BIN_DIR := bin/windows
endif
build:
	@echo "Building for $(OS)..."
	@mkdir -p $(BIN_DIR)
	@$(GO) build -o $(BIN_DIR)/$(BINARY_NAME)$(EXT) ./cmd/todo/

install: build
	@echo "Installing binary..."
	@cp $(BIN_DIR)/$(BINARY_NAME)$(EXT) /usr/local/bin/$(BINARY_NAME)$(EXT)

clean:
	go clean
