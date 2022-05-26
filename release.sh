#!/bin/sh

echo "Release version:"
read TARGET_VERSION

echo "Prepare release of cryptletter at $TARGET_VERSION"

git cliff --tag ${TARGET_VERSION} -o CHANGELOG.md && \
echo "Successfully generated changelog"

/bin/sh ./build.sh ${TARGET_VERSION}

git add ./web/package.json ./CHANGELOG.md
git commit -m "Release $TARGET_VERSION"
git tag -a ${TARGET_VERSION} -m "release.sh"

REPOSITORY="scribblerockerz/cryptletter"

docker build -t ${REPOSITORY}:latest -t ${REPOSITORY}:${TARGET_VERSION} . 2>&1 1>/dev/null && \
echo "Successfully build docker image for ${REPOSITORY}:${TARGET_VERSION}"

docker push ${REPOSITORY}:${TARGET_VERSION} && \
docker push ${REPOSITORY}:latest && \
echo "Successfully pushed image to hub.docker.com"

echo "Finished build"
