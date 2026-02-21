.PHONY: all build test clean generate

PROGRAM := gostub
CMD := $(CURDIR)/cmd
BIN := $(CURDIR)/bin
LD := "-extldflags '-static' -s -w"
CCFLAGS := CGO_ENABLED=0

all: generate test build

build: generate $(CMD) | $(BIN)
	@$(CCFLAGS) go build -ldflags=$(LD) -o $(BIN)/$(PROGRAM) $(CMD)/**
	@cp $(CURDIR)/config.example.yml $(BIN)/config.yaml
	@mkdir -p $(BIN)/stubs
	@touch $(BIN)/stubs/.placeholder
	@chmod +x $(BIN)/$(PROGRAM)

test: generate $(CMD)
	@go test ./...

clean:
	@rm -rf $(BIN)

generate:
	@go generate ./...

$(BIN):
	@mkdir -p $(BIN)
