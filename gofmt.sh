#!/bin/sh

cd "$(dirname "$0")"
exec gofmt -l -w -e -s -- $(find -type f -name '*.go')
