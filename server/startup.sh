#!/bin/sh

make migrateup
make sqlc-gen
air -c .air.toml
