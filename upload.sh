#!/bin/sh

set -e

NAME="eu.gcr.io/birota-cloud/github-pr-resource:1.11"

make build

docker build --platform linux/amd64 -t $NAME .
docker push $NAME
