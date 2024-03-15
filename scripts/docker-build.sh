#!/bin/sh -x

set -e
set -o errexit -o nounset

docker build --tag sighupio/go-jsonschema:latest --file Dockerfile --target go-jsonschema .
