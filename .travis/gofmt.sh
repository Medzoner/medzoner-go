#!/bin/bash
set -ev

test -z "$(gofmt -d -s ./pkg | tee /dev/stderr)"
