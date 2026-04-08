## Guidelines for pull requests

- Write tests for any changes.
- Separate unrelated changes into multiple pull requests.
- For bigger changes, make sure you start a discussion first by creating an issue and explaining the intended change.
- Ensure the build is green before you open your PR. The Pipelines build won't run by default on a remote branch, so enable Pipelines.

## Build

* [Go](https://golang.org/dl/)
* To build binary run `make build` in shell, it will produce `custom-gcl` binary.
* To install use `INSTALL_DIR=~/bin bash ./install.sh`, it will put `custom-gcl` binary under `~/bin` directory assuming it is in your `PATH`.

## Releasing a New Version

1. All notable changes comming with the new version should be documented in [CHANGELOG.md](https://raw.githubusercontent.com/smeshkov/golangci-lint-plugins/master/CHANGELOG.md).
2. Run lint with `make lint`, make sure everything is passing.
3. Run tests with `make test`, make sure everything is passing.
4. Bump the `TAG` variable inside the `Makefile` to the desired version, 
5. Make sure the change has been reviewed, approved and associated PR has been merged.
6. Push and trigger new binary release on GitHub via `make release`.
