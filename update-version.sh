#!/bin/bash

# Bump version using semver
VERSION=$(semver bump patch -p)

# Update version in Go code
sed -i "s/const version = \".*\"/const version = \"$VERSION\"/" main.go

# Update version in GitHub Actions workflow
sed -i "s/tag: v.*$/tag: v$VERSION/" .github/workflows/build-release.yml