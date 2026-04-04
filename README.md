# golangci-lint-plugins

Custom linter plugins for [golangci-lint](https://golangci-lint.run/) v2 using the [module plugin system](https://golangci-lint.run/docs/plugins/module-plugins/).

## Available Plugins

| Plugin | Description |
|--------|-------------|
| `modernfor` | Enforces Go 1.22+ `for i := range N` syntax instead of `for i := 0; i < N; i++` |

## Releasing a New Version

1. Commit your changes and push to `master`.
2. Tag and push:
   ```bash
   make release TAG=v0.1.0
   ```
   The tag must be a valid [Go module version](https://go.dev/ref/mod#versions) (e.g. `v0.1.0`, `v1.2.3`).

## Integrating into Another Project

### Prerequisites

Install golangci-lint v2:
```bash
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
```

### Step 1: Create `.custom-gcl.yml`

In the root of your target project, create `.custom-gcl.yml`:

```yaml
version: v2.11.4
plugins:
  - module: 'github.com/smeshkov/golangci-lint-plugins'
    import: 'github.com/smeshkov/golangci-lint-plugins/modernfor/plugin'
    version: v0.1.0  # use the latest published tag
```

For local development (unpublished changes), use `path` instead of `version`:
```yaml
version: v2.11.4
plugins:
  - module: 'github.com/smeshkov/golangci-lint-plugins'
    import: 'github.com/smeshkov/golangci-lint-plugins/modernfor/plugin'
    path: /absolute/path/to/golangci-lint-plugins
```

### Step 2: Configure `.golangci.yml`

Add the plugin to your project's `.golangci.yml`:

```yaml
version: "2"

linters:
  enable:
    - modernfor
  settings:
    custom:
      modernfor:
        type: module
        description: Enforces Go 1.22 range over int
```

### Step 3: Build and Run

```bash
# Build the custom binary (reads .custom-gcl.yml)
golangci-lint custom

# Run the linter
./custom-gcl run ./...
```

The custom binary is a drop-in replacement for `golangci-lint` — all standard flags and linters still work alongside the custom plugins.

### GitHub Actions CI

There are two approaches: download a pre-built binary from releases (faster), or build from source.

#### Option A: Pre-built binary (recommended)

Pre-built binaries are published to [GitHub Releases](https://github.com/smeshkov/golangci-lint-plugins/releases) on every tag. No `.custom-gcl.yml` needed in the target project.

```yaml
name: Lint

on:
  push:
    branches: [master, main]
  pull_request:

env:
  CUSTOM_GCL_VERSION: v0.1.0  # pin to a release tag

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - name: Download custom-gcl
        run: |
          curl -sL "https://github.com/smeshkov/golangci-lint-plugins/releases/download/${CUSTOM_GCL_VERSION}/custom-gcl-linux-amd64.tar.gz" | tar xz
          chmod +x custom-gcl

      - run: ./custom-gcl run ./...
```

#### Option B: Build from source

Requires `.custom-gcl.yml` with the plugin version in the target project (see Step 1 above).

```yaml
name: Lint

on:
  push:
    branches: [master, main]
  pull_request:

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod

      - run: go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4

      - run: golangci-lint custom

      - run: ./custom-gcl run ./...
```
