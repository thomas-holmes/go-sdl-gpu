GO_BIN?=go
CC=$(GO_BIN)
GOBUILD=$(CC) build
GOPKG='github.com/thomas-holmes/go-sdl-gpu/gpu'
ROOT=$(shell git rev-parse --show-toplevel)

all: test build

build:
	$(GOBUILD) $(GOPKG)

clean:
	$(GO_BIN) clean -cache $(GOPKG)

test:
	$(GO_BIN) test ./...

