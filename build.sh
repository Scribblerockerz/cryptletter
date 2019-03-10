#!/bin/sh

# build cryptletter executable
go build -o bin/cryptletter src/*.go

echo "Successfully build executable to bin/cryptletter"
