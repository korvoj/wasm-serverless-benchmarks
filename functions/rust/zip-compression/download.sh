#!/bin/bash

mkdir -p /tmp/test-dump/zip-compression/sample-files
curl -L -o /tmp/test-dump/zip-compression/sample-files/cirros-0.6.0-x86_64-rootfs.img.gz https://github.com/cirros-dev/cirros/releases/download/0.6.0/cirros-0.6.0-x86_64-rootfs.img.gz
curl -L -o /tmp/test-dump/zip-compression/sample-files/cirros-0.5.2-aarch64-rootfs.img.gz https://github.com/cirros-dev/cirros/releases/download/0.5.2/cirros-0.5.2-aarch64-rootfs.img.gz
