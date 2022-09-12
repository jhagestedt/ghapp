#!/bin/bash -e

set -euo pipefail

OSS=(linux darwin)
ARCHS=(amd64 arm64)

go get
go mod download

for OS in "${OSS[@]}"; do
  for ARCH in "${ARCHS[@]}"; do
    CGO_ENABLED=0 \
    GOOS="${OS}" \
    GOARCH="${ARCH}" \
    go build \
    -o "./dist/ghapp_${OS}_${ARCH}" \
    -ldflags "-X main.version=${VERSION}"
  done
done

for ASSET in ./dist/*; do
  echo "${ASSET}"
done