SOURCE_FILES = $(shell find $(CURDIR) -type f -name '*.go')

default: help

bin: $(SOURCE_FILES) ## Build binaries
	go build -o $(CURDIR)/bin/greeter .
	go build -o $(CURDIR)/bin/greeter-english ./builtin/english
	go build -o $(CURDIR)/bin/greeter-french ./builtin/french

.PHONY: clean
clean: ## Clean source tree
	rm -r $(CURDIR)/bin

.PHONY: help
help: ## Display target information
	@echo "Valid targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
