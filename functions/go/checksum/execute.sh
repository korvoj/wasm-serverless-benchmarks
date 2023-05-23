#/bin/bash

set -eu
REGISTRY_USER=$1
REGISTRY_PASSWORD=$2
IMAGE_REPOSITORY=ghcr.io/korvoj/wasm-serverless-benchmarks
IMAGE_NAME=wasmedge/go-aot-checksum
IMAGE_TAG=v1.0.0
WASM_BINARY_PATH=checksum.aot.wasm

ctr image pull --user $REGISTRY_USER:$REGISTRY_PASSWORD $IMAGE_REPOSITORY/$IMAGE_NAME:$IMAGE_TAG
ctr run --rm --runc-binary crun \
  --runtime io.containerd.runc.v2 \
  --label module.wasm.image/variant=compat-smart \
  --mount type=bind,src=../../sample-files/checksum,dst=/context,options=rbind:rw \
  $IMAGE_REPOSITORY/$IMAGE_NAME:$IMAGE_TAG wasmedge-go-aot-checksum \
  /$WASM_BINARY_PATH /context/cirros-0.6.0-x86_64-disk.img
