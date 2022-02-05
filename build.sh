#/bin/bash
set -e
#
CGO_ENABLED=0 GOOS=linux  go build -ldflags="-s -w" cmd/todo/todo-api.go
#