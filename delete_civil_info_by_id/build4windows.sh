#!/bin/bash
mkdir -p ./bin/windows/
GOOS=windows GOARCH=amd64 go build ./deleteInfo.go 
mv ./deleteInfo.exe ./bin/windows