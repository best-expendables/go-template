#!/usr/bin/env bash
echo ">>>> gen swagger"
swag init -d cmd/http -g http.go -o ./build --parseDependency