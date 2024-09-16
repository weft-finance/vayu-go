#!/bin/bash

# Source gvm to make it available in the script
if [ -f "$HOME/.gvm/scripts/gvm" ]; then
    source "$HOME/.gvm/scripts/gvm"
else
    echo "gvm is not installed or not found. Please install gvm first."
    exit 1
fi

cleanup() {
    if [ -f go.mod.bak ]; then
        echo "Restoring original go.mod file"
        mv go.mod.bak go.mod
        echo "Deleting backup go.mod.bak file"
        rm -f go.mod.bak
    fi
}

trap cleanup EXIT

pushd "$(pwd)" > /dev/null

if [ -z "$1" ]; then
    echo "Usage: $0 <go-version>"
    popd > /dev/null
    exit 1
fi

GO_VERSION=$1

if gvm list | grep -q "$GO_VERSION"; then
    echo "Go version $GO_VERSION is already installed."
else
    echo "Go version $GO_VERSION is not installed. Installing..."
    gvm install go"$GO_VERSION" --binary
fi

gvm list

go -version

echo "Switching to Go version $GO_VERSION"
gvm use go"$GO_VERSION" --default

echo "Backing up go.mod"
cp go.mod go.mod.bak

echo "Temporarily updating go.mod to use Go $GO_VERSION"
go mod edit -go="$GO_VERSION"

echo "Tidying Go modules"
go mod tidy

echo "Building the project"
if go build -v ./...; then
    echo "Build successful!"
else
    echo "Build failed!"
    popd > /dev/null
    exit 1
fi

echo "Running tests"
if go test -v ./...; then
    echo "All tests passed!"
else
    echo "Some tests failed!"
    popd > /dev/null
    exit 1
fi

popd > /dev/null
