#!/bin/bash

#rm -f jisho jisho.go
#go build -o jisho main.go

# gox -output="build/{{.OS}}/{{.Arch}}/{{.Dir}}" -osarch="linux/amd64 darwin/amd64 windows/amd64" -verbose
CGO_ENABLED=1 gox -output="build/" -osarch="linux/amd64" -verbose
