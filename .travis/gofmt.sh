#!/bin/bash
set -ev

test -z "$(gofmt ./pkg | tee /dev/stderr)"
