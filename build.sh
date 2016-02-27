#!/bin/bash

#rm -f jisho jisho.go
#go build -o jisho main.go

CGO_ENABLED=1 gox -output="build/{{.OS}}/{{.Arch}}/{{.Dir}}" -osarch="linux/amd64 darwin/amd64 windows/amd64"
