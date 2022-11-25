#!/bin/bash

VERSION="v0.1.1"

# Build multi-arch binaries
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/linux_amd64/jeb
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o build/linux_arm64/jeb
GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o build/linux_arm/jeb
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o build/darwin_amd64/jeb
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o build/darwin_arm64/jeb
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o build/windows_amd64/jeb

# Compress
tar -czvf jeb-linux-amd64-$VERSION.tgz -C build/linux_amd64 .
tar -czvf jeb-linux-arm64-$VERSION.tgz -C build/linux_arm64 .
tar -czvf jeb-linux-arm-$VERSION.tgz -C build/linux_arm .
tar -czvf jeb-darwin-amd64-$VERSION.tgz -C build/darwin_amd64 .
tar -czvf jeb-darwin-arm64-$VERSION.tgz -C build/darwin_arm64 .
tar -czvf jeb-windows-amd64-$VERSION.tgz -C build/windows_amd64 .

# Get sha1sum from binaries
mv build/*/*.tgz .
sha1sum *.tgz > sha1sums.txt
