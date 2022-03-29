#!/usr/bin/env bash

root=$1

cd "${root}"/cmd/migrate || exit

go run main.go --conf "${root}"/configs