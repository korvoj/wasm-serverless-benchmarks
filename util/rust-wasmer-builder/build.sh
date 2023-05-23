#!/bin/bash
set -eu

IMAGE_REPOSITORY=ghcr.io/korvoj/wasm-serverless-benchmarks
IMAGE_NAME=util/rust-wasmer-builder
TAG=1.64.0-bullseye

docker build -t $IMAGE_REPOSITORY/$IMAGE_NAME:$TAG .
docker push $IMAGE_REPOSITORY/$IMAGE_NAME:$TAG
