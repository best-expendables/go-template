#!/usr/bin/env bash
set -e

#getent hosts $POSTGRES_HOST >> /etc/hosts
#getent hosts $SENTINEL_HOST >> /etc/hosts

cd /go/src/bitbucket.org/gank-global/${PROJ_NAME}

export GO111MODULE=on

echo ">> Precompile dependencies"
go install "/go/src/bitbucket.org/gank-global/${PROJ_NAME}/cmd/http"

echo '>> Build migrations...'
go build -o ./build/migrate cmd/migrate/migrate.go
echo '>> Run migrations...'
./build/migrate up

echo '>> Build seeder...'
go build -o ./build/seeder cmd/seeder/seeder.go
# Dev has to seed data manually
echo '>> Run seeder...'
./build/seeder

# Reload app
/usr/local/bin/app/reload.sh "bitbucket.org/gank-global/${PROJ_NAME}/cmd/http" "/tmp/http"

echo "running"
