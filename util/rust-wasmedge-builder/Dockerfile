FROM rust:1.64.0-bullseye

RUN apt-get update && apt-get install -y curl && apt-get clean
RUN rustup target add wasm32-wasi
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -e all -p /usr/local
