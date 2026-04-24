#!/bin/bash
set -e

# Install Go
GO_VERSION="1.26.1"
curl -sL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" | tar -xz -C $HOME
export PATH="$HOME/go/bin:$PATH"

# Build WASM
rm -rf bin/web
mkdir -p bin/web
env GOOS=js GOARCH=wasm go build -o bin/web/astronaut.wasm github.com/bhaeussermann/ebitengine-astronaut/src
cp web/*.* bin/web
mkdir bin/web/images
cp images/*.* bin/web/images
