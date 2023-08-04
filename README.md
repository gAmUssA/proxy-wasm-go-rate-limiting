# proxy-wasm-go-rate-limiting

A prototype implementation of a rate-limiting filter written in Go,
using the proxy-wasm API for running on WebAssembly-enabled gateways.

## What's implemented

* "local" policy only, using the SHM-based key-value store

## What's missing

* Getting proper route and service ids for producing identifiers.
* Other policies, which would require additional features from the
  underlying system, such as calling out a Redis instance.

## Build requirements

* [tinygo](https://tinygo.org)
* [proxy-wasm-go-sdk](github.com/tetratelabs/proxy-wasm-go-sdk)
* [ffjson](https://github.com/pquerna/ffjson)

## Building and running

Once the environment is set up with `tinygo` and `ffjson` in your PATH,
build the filter running `make` and `make test` to run integration test.

You also need the `kong_wasm_rate_limiting_counters` shared memory
key-value store enabled in your Kong configuration. 
One way to achieve this is via the following environment variable:

```sh
export KONG_NGINX_WASM_SHM_KONG_WASM_RATE_LIMITING_COUNTERS=12m
```
