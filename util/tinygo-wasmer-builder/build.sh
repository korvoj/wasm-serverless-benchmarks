#!/bin/bash
set -eu

IMAGE_REPOSITORY=ghcr.io/korvoj/wasm-serverless-benchmarks
IMAGE_NAME=util/tinygo-wasmer-builder
TAG=0.26.0

docker build -t $IMAGE_REPOSITORY/$IMAGE_NAME:$TAG .
docker push $IMAGE_REPOSITORY/$IMAGE_NAME:$TAG
