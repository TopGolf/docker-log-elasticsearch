#!/bin/bash

set -e

# fetch all tags from the remote
git fetch --tags

# find last tag
LAST_TAG="$(git describe --tags "$(git rev-list --tags --max-count=1)" 2>/dev/null)"

# create a new tag version
NEW_TAG="$(awk -F. '{printf "%d.%d.%d", $1, $2, $3+1}' <(echo "$LAST_TAG"))"

export PLUGIN_TAG="$NEW_TAG"
make

docker login -u "$DOCKER_USERNAME" -p "$DOCKER_PASSWORD" &>/dev/null
make push

git log --oneline "${LAST_TAG}..HEAD"

git tag -a "$NEW_TAG" -m "$(git log --oneline ${LAST_TAG}..HEAD)"

git tag -n10
