#!/bin/sh

make migrate-up

exec "$@"
