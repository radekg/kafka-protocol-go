# Kafka Protocol for go and tinygo

This is a fork of the Kafka protocol implementation from [grepplabs/kafka-proxy](https://github.com/grepplabs/kafka-proxy). This fantastic work is a great donor because it:

- Does not use reflection.
- Does not use goroutines.
- Does not use channels.

Which results in:

- Being able to compile with [tinygo-org/tinygo](https://github.com/tinygo-org/tinygo).
- Because it can be compiled with _tinygo_, it can be compiled with [tetratelabs/proxy-wasm-go-sdk](https://github.com/tetratelabs/proxy-wasm-go-sdk).
  - Which means that it can be used as a foundation for Envoy WebAssembly Layer 7 proxies.

## Why forking?

The upstream Kafka Proxy protocol implements sufficient encoding and decoding to support its own functionality, but does not implement everything else required to parse messages beyond request headers, response headers, `Metadata`, and `FindCoordinator`. The maintainer of Kafka Proxy isn't keen on supporting anything in the protocol encoder/decoder that does not find a direct use in the upstream project.

The goal of this fork is to add support for all Apache Kafka messages.

## License 

Apache 2.0.

## Run tests

```sh
go test -count=1 -v ./...
```
