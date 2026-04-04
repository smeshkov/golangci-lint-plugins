TAG ?= 0.0.2
BINARY_NAME := custom-gcl
GOLANGCI_LINT_VERSION := v2.11.4

.PHONY: build clean tag release

build:
	golangci-lint custom --version $(GOLANGCI_LINT_VERSION) --name build/$(BINARY_NAME)

clean:
	rm -rf build/*

# Usage: make tag V=v0.1.0
tag:
	@if [ -z "$(TAG)" ]; then echo "Usage: make tag V=v0.1.0"; exit 1; fi
	git tag -a $(TAG) -m "Release $(V)"
	@echo "Tagged $(TAG). Run 'git push origin $(TAG)' to publish."

# Usage: make release V=v0.1.0
release: tag
	git push origin $(TAG)
