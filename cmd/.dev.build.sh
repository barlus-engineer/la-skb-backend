#!/bin/bash

# Build the Go project
go build -o cmd/laskb_server cmd/main.go

if [ -d "out" ]; then
    rm -r out
fi
mkdir out
mv cmd/laskb_server out/laskb
cp -r locale out/locale

if [ $? -eq 0 ]; then
    echo "Build successful!"
else
    echo "Build failed!"
    exit 1
fi