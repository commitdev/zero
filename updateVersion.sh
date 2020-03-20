#!/usr/bin/env bash

VERSION=$1

if [ -z "$VERSION" ]; then
    echo "Usage: ./updateVersion.sh <version>"
    exit 1
fi

sed "s/\t\tfmt.Println(\"v0.0.0\")/\t\tfmt.Println(\"${VERSION}\")/g" cmd/version.go > cmd/version.go.update

mv cmd/version.go.update cmd/version.go