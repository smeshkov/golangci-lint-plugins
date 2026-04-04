#!/bin/bash
set -euo pipefail

REPO="smeshkov/golangci-lint-plugins"
BINARY_NAME="custom-gcl"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Detect OS and architecture
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH" >&2; exit 1 ;;
esac

# Get latest version or use specified version
VERSION="${VERSION:-$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | cut -d '"' -f 4)}"
if [ -z "$VERSION" ]; then
  echo "Failed to determine latest version" >&2
  exit 1
fi

TARBALL="${BINARY_NAME}-${OS}-${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${TARBALL}"

echo "Installing ${BINARY_NAME} ${VERSION} (${OS}/${ARCH})..."

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

curl -sL "$URL" -o "${TMP}/${TARBALL}"
tar -xzf "${TMP}/${TARBALL}" -C "$TMP"

install -m 755 "${TMP}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
echo "Installed ${BINARY_NAME} to ${INSTALL_DIR}/${BINARY_NAME}"
