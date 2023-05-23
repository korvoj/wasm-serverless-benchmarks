#!/bin/bash

wget https://static.mrezhi.net/wasm-serverless-benchmarks/sample-files.zip
unzip sample-files.zip -d sample-files
rm sample-files.zip
mv sample-files/Benchmarks/* sample-files
rm -r sample-files/Benchmarks
