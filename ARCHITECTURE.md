# Architecture

Each linter lives in its own package under the repo root (e.g. `modernfor/`):

- `modernfor/modernfor.go` — Analyzer implementation (AST inspection logic)
- `modernfor/plugin/plugin.go` — Plugin entry point wrapping the analyzer for golangci-lint

Key config files:
- `.custom-gcl.yml` — Defines which plugins to compile into the custom binary; uses `path` for local builds
- `.golangci.yml` — Enables the custom linters (type: module) so the built binary runs them