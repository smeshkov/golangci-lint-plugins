# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/).

## [v0.0.6] - 2026-04-08

### Added

- `useany` plugin: enforces using `any` instead of `interface{}`.
- GitHub Actions release workflow for pre-built binaries (linux/darwin, amd64/arm64).

## [v0.0.5] - 2026-04-05

### Added

- `modernfor` plugin: enforces Go 1.22+ `for i := range N` syntax instead of traditional `for i := 0; i < N; i++` loops.
- GitHub Actions release workflow for pre-built binaries (linux/darwin, amd64/arm64).
- Integration guide and CI examples in README.
