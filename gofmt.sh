#!/bin/sh

cd "$(dirname "$0")"
exec gofmt -l -w -e -s -tabs $(find -type f -name '*.go')
