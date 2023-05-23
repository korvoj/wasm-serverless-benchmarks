FROM tinygo/tinygo:0.26.0

USER root
RUN apt-get update && apt-get install -y curl libtinfo5 && apt-get clean
RUN curl -L -o /tmp/wasmer-2.3.0.tar.gz https://github.com/wasmerio/wasmer/releases/download/2.3.0/wasmer-linux-amd64.tar.gz
RUN mkdir /tmp/wasmer && tar -xvf /tmp/wasmer-2.3.0.tar.gz -C /tmp/wasmer && mv /tmp/wasmer/bin/wasmer /usr/local/bin/wasmer && chmod +x /usr/local/bin/wasmer && rm -r /tmp/wasmer-2.3.0.tar.gz /tmp/wasmer
WORKDIR /home/tinygo

USER tinygo
