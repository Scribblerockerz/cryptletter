#!/bin/sh

# Turorial
# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

# build cryptletter executable
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/cryptletter src/*.go && \
echo "Successfully build executable to ./bin/cryptletter"

npm run build 2>&1 1>/dev/null && \
echo "Successfully build assets to ./public"

docker build -t scribblerockerz/cryptletter:$(git tag -l --points-at HEAD) . 2>&1 1>/dev/null && \
echo "Successfully build docker image for scribblerockerz/cryptletter"

docker push scribblerockerz/cryptletter && \
echo "Successfully pushed image to hub.docker.com"

echo "Finished build"

