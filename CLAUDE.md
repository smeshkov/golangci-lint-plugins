# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Custom golangci-lint plugins (module plugin system). Each plugin is an `analysis.Analyzer` from `golang.org/x/tools/go/analysis`, exposed via a `plugin/plugin.go` wrapper that exports a `New(conf any) ([]*analysis.Analyzer, error)` function — the signature golangci-lint expects.

## Development process

- `make checkfmt`, `make lint` and `make test` must pass.
- always make sure to add tests for any new code you write.
- tests should be placed next to the code they test.

## Build Commands

```bash
make build           # Build custom golangci-lint binary to build/custom-gcl
make clean           # Remove build artifacts
make tag TAG=0.1.0   # Create annotated git tag
make release TAG=0.1.0   # Tag and push to remote
```

Requires `golangci-lint` v2 installed (`go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest`).

## Architecture

See `ARCHITECTURE.md`.

## Adding a New Linter

1. Create `<name>/` package with an exported `Analyzer` variable
2. Create `<name>/plugin/plugin.go` with the `New` function returning the analyzer
3. Add the plugin to `.custom-gcl.yml` and enable it in `.golangci.yml`
