#!/bin/bash
go install github.com/githubnemo/CompileDaemon@v1.4.0
export GO_ENABLED=0 GOOS=linux
CompileDaemon --build="go build -v -o server" --command=/app/src/server