#!/bin/bash
set -ev

gofmt -w pkg
test -z "$(gofmt -w -d -s ./pkg | tee /dev/stderr)"
