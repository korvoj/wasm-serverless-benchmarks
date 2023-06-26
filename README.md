# WebAssembly as an Enabler for Next Generation Serverless Computing

## Introduction
This repository contains complementary data for the paper titled "WebAssembly as an Enabler for Next Generation Serverless Computing" (*Kjorveziroski, V., Filiposka, S. WebAssembly as an Enabler for Next Generation Serverless Computing. J Grid Computing 21, 34 (2023). [https://doi.org/10.1007/s10723-023-09669-8](https://doi.org/10.1007/s10723-023-09669-8)*).

The repository is organized around 5 core directories:

- `functions` – contains the source code of the serverless benchmarking functions used for evaluating the performance of container and WebAssembly runtime environments. Dockerfiles are present for building specific OCI images targeting: distroless  base images, debian base images, WasmEdge, Wasmer, and Wasmtime.
- `results` – contains the performance results of the 13 different benchmarking functions, showcasing their cold start and execution performance. Each executed function using a given runtime is represented by a single `.txt` file, containing the execution times of each run on a separate run (in seconds). All functions are run exactly 100 times.
- `run` – contains shell scripts which allow semi-automatic deployment and instantiation of the benchmarking runs.
- `sample-files` – empty by default, can be populated by executing the `download-sample-files.sh` shell script which fetches the required input files for each of the benchmarks and creates the necessary folder structure for the output results. A separate download script is required to overcome the file size limitations imposed by certain source code management platforms.
- `util` – contains the Dockerfiles for the builder Docker images, used for compiling the source code of the benchmarking functions.

A Jupyter notebook with a step-by-step analysis off all of the obtained results is also available in the repository, `jupyter-wasm-serverless-benchmarks.ipynb`. Any public or private Jupyter notebooks server can be used for exploring the contents of the file.

## Running the Benchmarks

To run the benchmarks locally, the following software is required:

- containerd
- crun
- wasmedge
- wasmtime
- wasmer
- buildah
- docker

Please note that `crun` supports integration with only a single WebAssembly runtime at a given point in time.

The steps for reproducing the benchmarks are:

1. Clone this repository
2. Choose which runtime will be tested first ([WasmEdge](#configuring-crun-for-wasmedge-support), [Wasmtime](#configuring-crun-for-wasmtime-support), or [Wasmer](#configuring-crun-for-wasmer-support)) and follow the provided instructions for installing the prerequisites.
3. In case `AOT` execution needs to be evaluated for either Wasmtime or Wasmer make sure to make the required changes to the `crun` source code before compiling, as discussed in the relevant runtime subsections below.
4. Build the OCI images from scratch by navigating to the `functions/go` and `functions/rust` subfolders and manually editing the `build.sh` script files in order to provide the necessary information for the OCI image registry, before executing the scripts.
5. Download the sample files which are required for the execution of some functions by running the `download-sample-files.sh` script.
6. Run the tests. Each test has a dedicated script located in the `run/` subfolder. Please make sure to edit the OCI image registry location and credentials before executing the run scripts.
7. The results from each test will be available in the `results/` subfolder. Launch the Jupyter notebook `jupyter-wasm-serverless-benchmarks.ipynb` to analyze the results or repeat steps 2-6 to evaluate the performance of a different runtime.

### Configuring Crun for WasmEdge support

The following set of commands represents the minimum viable configuration for running the benchmarking functions in a WasmEdge environment:

```bash
# install necessary dependencies
apt install make git gcc build-essential pkgconf libtool libsystemd-dev libprotobuf-c-dev libcap-dev libseccomp-dev libyajl-dev go-md2man libtool autoconf python3 automake

# install wasmedge
curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -e all -p /usr/local

# clone crun repository
cd /opt
git clone https://github.com/containers/crun
cd crun
git checkout 1.6

# compile crun with wasmedge support
./autogen.sh
./configure --with-wasmedge
make
make install

# install and configure containerd
apt install containerd
mkdir -p /etc/containerd/
cd /etc/containerd
bash -c "containerd config default > /etc/containerd/config.toml"
wget https://raw.githubusercontent.com/second-state/wasmedge-containers-examples/main/containerd/containerd_config.diff
patch -d/ -p0 < containerd_config.diff
systemctl restart containerd
```

### Configuring Crun for Wasmer support

The following set of commands represents the minimum viable configuration for running the benchmarking functions in a Wasmer environment:

```bash
# install necessary dependencies
apt install make git gcc build-essential pkgconf libtool libsystemd-dev libprotobuf-c-dev libcap-dev libseccomp-dev libyajl-dev go-md2man libtool autoconf python3 automake

# install wasmer (2.3.0)
curl https://get.wasmer.io -sSfL | sh
source ~/.wasmer/wasmer.sh
apt install libtinfo5

# clone crun repository
cd /opt
git clone https://github.com/containers/crun
cd crun
git checkout 1.6

# copy wasmer libraries and headers to required locations
cp ~/.wasmer/lib/libwa* /usr/lib/
cp ~/.wasmer/include/w* /usr/include/

# compile crun with wasmer support
./autogen.sh
./configure --with-wasmer
make
make install

# install and configure containerd
apt install containerd
mkdir -p /etc/containerd/
cd /etc/containerd
bash -c "containerd config default > /etc/containerd/config.toml"
wget https://raw.githubusercontent.com/second-state/wasmedge-containers-examples/main/containerd/containerd_config.diff
patch -d/ -p0 < containerd_config.diff
systemctl restart containerd
```

Note that the `crun` integration with Wasmer only supports JIT compiled WebAssembly modules by default. To add support for AOT compilation, the `src/libcrun/handlers/wasmer.c` file needs to be manually edited, and all occurrences of `wasm_module_new` need to be replaced with `wasm_module_deserialize`. Once the changes are made, `crun` needs to be recompiled.  Further documentation about the `wasm_module_deserialize()` method is available on [Wasmer's official documentation page]( https://docs.rs/wasmer/latest/wasmer/struct.Module.html#method.deserialize).

### Configuring Crun for Wasmtime support

The following set of commands represents the minimum viable configuration for running the benchmarking functions in a Wasmtime environment:

```bash
# install necessary dependencies
apt install make git gcc build-essential pkgconf libtool libsystemd-dev libprotobuf-c-dev libcap-dev libseccomp-dev libyajl-dev go-md2man libtool autoconf python3 automake

# install wasmtime (wasmtime 1.0.1)
curl https://wasmtime.dev/install.sh -sSf | bash

# clone crun repository
cd /opt
git clone https://github.com/containers/crun
cd crun
git checkout 1.6

# download wasmtime headers from https://github.com/bytecodealliance/wasmtime
wget https://github.com/bytecodealliance/wasmtime/releases/download/v1.0.1/wasmtime-v1.0.1-x86_64-linux-c-api.tar.xz

# copy wasmtime libraries and headers to required locations
cp -r wasmtime-v1.0.1-x86_64-linux-c-api/include/* /usr/include/
cp -r wasmtime-v1.0.1-x86_64-linux-c-api/lib/* /usr/lib

# compile crun with wasmtime support
./autogen
./configure --with-wasmtime
make
make install

# install and configure containerd
apt install containerd
mkdir -p /etc/containerd/
cd /etc/containerd
bash -c "containerd config default > /etc/containerd/config.toml"
wget https://raw.githubusercontent.com/second-state/wasmedge-containers-examples/main/containerd/containerd_config.diff
patch -d/ -p0 < containerd_config.diff
systemctl restart containerd
```

Note that the `crun` integration with Wasmtime only supports JIT compiled WebAssembly modules by default. To add support for AOT compilation, the `src/libcrun/handlers/wasmtime.c` file needs to be manually edited, and all occurrences of `wasmtime_module_new` need to be replaced with `wasmtime_module_deserialize`. Once the changes are made, `crun` needs to be recompiled. Further documentation about these two methods is available on the Wasmtime official documentation page ([wasmtime_module_deserialize()](https://docs.wasmtime.dev/c-api/module_8h.html), [wasmtime_module_new()](https://docs.wasmtime.dev/c-api/module_8h.html#aa9ce0ce436f27b54ccacf044d346776a)).