#!/bin/bash

if [ "$1" == "" ]; then
    VERSION=$(git describe --tags $(git rev-list --tags --max-count=1))
    echo "Building with latest git version $VERSION"
else
    VERSION="$1"
    echo "Building with explicit version $VERSION"
fi

# build frontend app
cd web && \
pnpm version $VERSION --allow-same-version 2>&1 1>/dev/null && \
pnpm run build 2>&1 1>/dev/null && \
cd .. &&  \
echo "Successfully build assets to ./web/dist"

LDFLAGS="-X github.com/Scribblerockerz/cryptletter/cmd/cryptletter.Version=$VERSION"

# build executables
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="$LDFLAGS" -o bin/cryptletter *.go && \
echo "Successfully build executable for LINUX to ./bin/cryptletter"

CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags="$LDFLAGS" -o bin/cryptletter-macos *.go && \
echo "Successfully build executable for MACOS to ./bin/cryptletter-macos"
