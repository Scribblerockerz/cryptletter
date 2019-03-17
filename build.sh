#!/bin/sh

# Turorial
# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

# build cryptletter executable
go build -o bin/cryptletter src/*.go

echo "Successfully build executable to bin/cryptletter"
