#!/usr/bin/env bash

getent hosts $POSTGRES_HOST >> /etc/hosts
getent hosts $SENTINEL_HOST >> /etc/hosts

echo "Starting application"
./build/http