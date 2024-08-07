#!/bin/sh
./.github/gofmt.sh
go vet ./...
go test -v -cover -coverpkg=./pkg/... -covermode=count -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
rm coverage.out
find ./pkg -type f -name '*.go' -exec wc -l {} +
find ./pkg -type f -name '*.go' -exec sed '/^\s*$/d' {} + | wc -l; echo ' ↳ total (non-blank) lines of code'
