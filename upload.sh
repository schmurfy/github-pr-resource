#!/bin/sh

set -e

NAME="eu.gcr.io/birota-cloud/github-pr-resource:1.9"

make build

docker build -t $NAME .
docker push $NAME
