#/bin/bash

set -eu
REGISTRY_USER=$1
REGISTRY_PASSWORD=$2
IMAGE_REPOSITORY=ghcr.io/korvoj/wasm-serverless-benchmarks
MODE=$3
IMAGE_TAG=v1.0.0

case $MODE in
  debian)
    echo -n "Executing as a debian image..."
    IMAGE_NAME=debian/rust-zip-compression
    WASM_BINARY_PATH=zip-compression
    RUN_NAME=debian-rust-zip-compression
    ;;
  distroless)
    echo -n "Executing as a distroless image..."
    IMAGE_NAME=distroless/rust-zip-compression
    WASM_BINARY_PATH=zip-compression
    RUN_NAME=distroless-rust-zip-compression
    ;;
  wasmedge-jit)
    echo -n "Executing with WasmEdge JIT..."
    IMAGE_NAME=wasmedge/rust-jit-zip-compression
    WASM_BINARY_PATH=zip-compression.wasm
    RUN_NAME=wasmedge-rust-jit-zip-compression
    ;;
  wasmedge-aot)
    echo -n "Executing with WasmEdge..."
    IMAGE_NAME=wasmedge/rust-aot-zip-compression
    WASM_BINARY_PATH=zip-compression.aot.wasm
    RUN_NAME=wasmedge-rust-aot-zip-compression
    ;;
  wasmtime-jit)
    echo "Executing with WasmTime JIT..."
    IMAGE_NAME=wasmtime/rust-jit-zip-compression
    WASM_BINARY_PATH=zip-compression.wasm
    RUN_NAME=wasmtime-rust-jit-zip-compression
    ;;
  wasmtime-aot)
    echo "Executing with WasmTime AOT..."
    IMAGE_NAME=wasmtime/rust-aot-zip-compression
    WASM_BINARY_PATH=zip-compression.aot.wasm
    RUN_NAME=wasmtime-rust-aot-zip-compression
    ;;
  wasmer-jit)
    echo "Executing with Wasmer JIT..."
    IMAGE_NAME=wasmer/rust-jit-zip-compression
    WASM_BINARY_PATH=zip-compression.wasm
    RUN_NAME=wasmer-rust-jit-zip-compression
    ;;
  wasmer-aot)
    echo "Executing with Wasmer AOT..."
    IMAGE_NAME=wasmer/rust-aot-zip-compression
    WASM_BINARY_PATH=zip-compression.aot.wasm
    RUN_NAME=wasmer-rust-aot-zip-compression
    ;;
  *)
    echo "Unsupported runtime. Please choose one from the following: wasmedge-jit, wasmedge-aot, wasmtime-jit, wasmtime-aot, wasmer-jit, wasmer-aot"
    exit -1
    ;;
esac


ctr image pull --user $REGISTRY_USER:$REGISTRY_PASSWORD $IMAGE_REPOSITORY/$IMAGE_NAME:$IMAGE_TAG
ctr run --rm --runc-binary crun \
  --runtime io.containerd.runc.v2 \
  --label module.wasm.image/variant=compat-smart \
  --mount type=bind,src=../../sample-files/zip-compression,dst=/context,options=rbind:rw \
  $IMAGE_REPOSITORY/$IMAGE_NAME:$IMAGE_TAG $RUN_NAME \
  /$WASM_BINARY_PATH /context/sample-files /context/sample-files.zip
