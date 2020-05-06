#!/bin/bash

GOPATH="$GOROOT:$(pwd)/git/"
export GOPATH

echo "Install GoLint"
go get -u golang.org/x/lint/golint
command -v golint || exit 1

echo "Install GoDoc"
go get golang.org/x/tools/cmd/godoc
command -v godoc || exit 1

echo "Install yamllint"
brew install yamllint
command -v yamllint || exit 1

echo "$0 done"
