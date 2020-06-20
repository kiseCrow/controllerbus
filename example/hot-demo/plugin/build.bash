#!/bin/bash
set -eo pipefail

export GO111MODULE=on
export $(go env | grep GOOS)
GOOS="${GOOS%\"}"
GOOS="${GOOS#\"}"
if [[ "$GOOS" != "linux" ]]; then
    echo "This only works on GOOS=linux."
    exit 1
fi

echo "Building plugin binary..."
go build -o plugin-bin.so -v -buildmode=plugin ./plugin-bin

echo "Compiled plugin-bin.so successfully."

