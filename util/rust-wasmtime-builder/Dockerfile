FROM rust:1.64.0-bullseye

RUN apt-get update && apt-get install -y curl && apt-get clean
RUN rustup target add wasm32-wasi
RUN curl https://wasmtime.dev/install.sh -sSf | bash
RUN ln -s /root/.wasmtime/bin/wasmtime /usr/local/bin/wasmtime
