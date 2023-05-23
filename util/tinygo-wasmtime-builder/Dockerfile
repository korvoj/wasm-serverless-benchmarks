FROM tinygo/tinygo:0.26.0

USER root
RUN apt-get update && apt-get install -y curl xz-utils && apt-get clean
RUN curl -L -o /tmp/wasmtime-v1.0.1.tar.xz https://github.com/bytecodealliance/wasmtime/releases/download/v1.0.1/wasmtime-v1.0.1-x86_64-linux.tar.xz
RUN mkdir /tmp/wasmtime && tar -xvf /tmp/wasmtime-v1.0.1.tar.xz -C /tmp/wasmtime && mv /tmp/wasmtime/wasmtime-v1.0.1*/wasmtime /usr/local/bin/wasmtime && chmod +x /usr/local/bin/wasmtime && rm -r /tmp/wasmtime-v1.0.1.tar.xz /tmp/wasmtime

USER tinygo
