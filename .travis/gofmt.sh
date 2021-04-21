#!/bin/bash
set -ev

test -z "$(gofmt -w -d -s ./pkg | tee /dev/stderr)"
