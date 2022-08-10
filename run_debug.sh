#!/usr/bin/env bash

go build -gcflags="all=-N -l" -o ./bin/app ./cmd/app
go install github.com/go-delve/delve/cmd/dlv@master
dlv debug --headless --listen=:4000 --only-same-user=false --api-version=2 --accept-multiclient --log --log-output=rpc ./cmd/app/main.go
