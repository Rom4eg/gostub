.PHONY: all build test clean generate

PROGRAM := gostub
CMD := $(CURDIR)/cmd/gostub
BIN := $(CURDIR)/bin
LD := "-extldflags '-static' -s -w"
CCFLAGS := GOOS=linux GOARCH=amd64 CGO_ENABLED=0

export EXAMPLE_STUB

build: $(CMD) | $(BIN)
	@$(CCFLAGS) go build -ldflags=$(LD) -o $(BIN)/$(PROGRAM) $(CMD)/**
	@cp $(CURDIR)/config.example.yml $(BIN)/config.yaml
	@mkdir -p $(BIN)/stubs
	@touch $(BIN)/stubs/.placeholder
	@mkdir $(BIN)/stubs/default
	@touch $(BIN)/stubs/default/index
	@echo "$$EXAMPLE_STUB" > $(BIN)/stubs/default/index
	@chmod +x $(BIN)/$(PROGRAM)

test: $(CMD)
	@go test ./...

clean:
	@rm -rf $(BIN)

generate:
	@go generate ./...

$(BIN):
	@mkdir -p $(BIN)


define EXAMPLE_STUB
{{- define "main" -}}
	{{- .SetCode 200 -}}
gostub is ready!
Check out the documentation: https://gostub.readthedocs.io/
Visit the project page: https://github.com/Rom4eg/gostub
{{ end }}
endef
