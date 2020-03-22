# bhyve-client

Simple client for the B-Hyve Smart Sprinkler

## Simple Client for Controlling B-Hyve Smart Sprinkler

## Status

![CI](https://github.com/campbellmcneill/bhyve-client/workflows/CI/badge.svg)

## Manual

Run `make help` to list available commands:

``` make
  λ  make help

 Choose a command run in b-hyve client:

  install   Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
  compile   Compile the binary.
  exec      Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
  run       Run the binary created as part of this project. Compiles first.
  clean     Clean build files. Runs `go clean` internally.
```
