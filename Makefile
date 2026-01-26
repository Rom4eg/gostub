.PHONY: all build test clean

PROGRAM := gostub
CMD := $(CURDIR)/cmd
BIN := $(CURDIR)/bin

$(BIN):
	@mkdir -p $(BIN)

all: test build

build: $(CMD) | $(BIN)
	@go build -o $(BIN)/$(PROGRAM) $(CMD)/**
	@cp $(CURDIR)/config.example.yml $(BIN)/config.yaml
	@mkdir -p $(BIN)/stubs

test: $(CMD)
	@go test ./...

clean:
	@rm -rf $(BIN)
