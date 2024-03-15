#!/bin/sh -x

set -e
set -o errexit -o nounset

docker build --tag sighupio/go-jsonschema:tools-latest --file Dockerfile --target tools .
