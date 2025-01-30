#!/bin/bash

set -e  # Exit if any command fails

BUILD_DIR="public"

# Ensure build directory exists
mkdir -p "$BUILD_DIR"

# Compile Go code to WebAssembly
echo "Compiling Go to WASM..."
GOOS=js GOARCH=wasm go build -o "$BUILD_DIR/main.wasm" main.go

# Copy the Go WebAssembly runtime
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" "$BUILD_DIR/"

echo "Build complete: public/index.html"
