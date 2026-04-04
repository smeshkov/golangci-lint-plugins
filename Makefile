TAG ?= 0.0.5
BINARY_NAME := custom-gcl
GOLANGCI_LINT_VERSION := v2.11.4

.PHONY: build clean test tag release

build:
	golangci-lint custom --version $(GOLANGCI_LINT_VERSION) --name build/$(BINARY_NAME)

clean:
	rm -rf build/*

test:
	go test ./...

# Usage: make tag V=v0.1.0
tag:
	@if [ -z "$(TAG)" ]; then echo "Usage: make tag V=v0.1.0"; exit 1; fi
	git tag -a v$(TAG) -m "Release v$(TAG)"
	@echo "Tagged v$(TAG). Run 'git push origin v$(TAG)' to publish."

# Usage: make release V=v0.1.0
release: tag
	git push origin v$(TAG)
