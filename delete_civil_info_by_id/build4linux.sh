#!/bin/bash
mkdir -p ./bin/linux/
GOOS=linux GOARCH=amd64 go build ./deleteInfo.go 
mv ./deleteInfo ./bin/linux