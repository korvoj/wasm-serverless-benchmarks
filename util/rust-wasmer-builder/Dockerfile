FROM rust:1.64.0-bullseye

RUN apt-get update && apt-get install -y curl libtinfo5 && apt-get clean
RUN rustup target add wasm32-wasi
RUN curl https://get.wasmer.io -sSfL | sh
RUN ln -s /root/.wasmer/bin/wasmer /usr/local/bin/wasmer
