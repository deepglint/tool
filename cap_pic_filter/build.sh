#!/bin/bash
mkdir -p release

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o release/cap_pic_filter.exe cap_pic_filter.go
