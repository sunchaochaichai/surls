#!/usr/bin/env bash

docker build \
-f docker/image_build/Dockerfile \
../ \
-t=surls-img

docker-compose \
-f docker/docker-compose.yaml \
up