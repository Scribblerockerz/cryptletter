#!/bin/sh

# Turorial
# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

# build frontend app
cd web && npm run build 2>&1 1>/dev/null && cd .. &&  \
echo "Successfully build assets to ./web/dist"

# build executables
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/cryptletter *.go && \
echo "Successfully build executable for LINUX to ./bin/cryptletter"

CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/cryptletter-macos *.go && \
echo "Successfully build executable for MACOS to ./bin/cryptletter-macos"


docker build -t scribblerockerz/cryptletter:$(git tag -l --points-at HEAD) . 2>&1 1>/dev/null && \
echo "Successfully build docker image for scribblerockerz/cryptletter"

docker push scribblerockerz/cryptletter && \
echo "Successfully pushed image to hub.docker.com"

echo "Finished build"

