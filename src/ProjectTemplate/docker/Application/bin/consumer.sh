#!/usr/bin/env bash
set -e

cd /go/src/bitbucket.org/gank-global/${PROJ_NAME}

export GO111MODULE=on

echo ">> Precompile dependencies"
go install -a "bitbucket.org/gank-global/${PROJ_NAME}/cmd/consumer"

/usr/local/bin/app/reload.sh "bitbucket.org/gank-global/${PROJ_NAME}/cmd/consumer" "/tmp/consumer"