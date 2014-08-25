#!/bin/sh

BASE="$(dirname "$0")"
exec gofmt -l -w -e -s -- $(find "$BASE" -type f -name '*.go')
