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

REPOSITORY="scribblerockerz/cryptletter"
VERSION=$(git tag -l --points-at HEAD)
docker build -t ${REPOSITORY}:latest -t ${REPOSITORY}:${VERSION} . 2>&1 1>/dev/null && \
echo "Successfully build docker image for ${REPOSITORY}:${VERSION}"

docker push ${REPOSITORY}:${VERSION} && \
docker push ${REPOSITORY}:latest && \
echo "Successfully pushed image to hub.docker.com"

echo "Finished build"

