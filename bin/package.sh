#!/bin/bash -e

set -euo pipefail

docker buildx build . \
--platform=linux/amd64,linux/arm64
