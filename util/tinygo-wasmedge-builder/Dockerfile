FROM tinygo/tinygo:0.26.0

USER root
RUN apt-get update && apt-get install -y curl && apt-get clean
RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -e all -p /usr/local
USER tinygo
