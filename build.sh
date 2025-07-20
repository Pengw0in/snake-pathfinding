#!/bin/bash

# Create builds directory
mkdir -p builds

echo "Building Snake Game for multiple platforms..."

# Build for Windows
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o builds/snake-game-windows-amd64.exe ./src

echo "Building for Windows (386)..."
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o builds/snake-game-windows-386.exe ./src

# Build for macOS
echo "Building for macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o builds/snake-game-macos-arm ./src

echo "Build complete! Check the 'builds' directory."
ls -la builds/
