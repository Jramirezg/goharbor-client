# goharbor-client

[![GitHub license](https://img.shields.io/github/license/mittwald/goharbor-client.svg?style=flat-square)](https://github.com/mittwald/goharbor-client/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/mittwald/goharbor-client?style=flat-square)](https://goreportcard.com/badge/github.com/mittwald/goharbor-client)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/mittwald/goharbor-client/v3)
[![Release](https://img.shields.io/github/release/mittwald/goharbor-client.svg?style=flat-square)](https://github.com/mittwald/goharbor-client/releases/latest)

[![Maintainability](https://api.codeclimate.com/v1/badges/a765bafaa29f6f8fdde7/maintainability)](https://codeclimate.com/github/mittwald/goharbor-client/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a765bafaa29f6f8fdde7/test_coverage)](https://codeclimate.com/github/mittwald/goharbor-client/test_coverage)
[![Actions Status](https://github.com/mittwald/goharbor-client/workflows/Test/badge.svg)](https://github.com/mittwald/goharbor-client/actions)
---

:warning: **STABILITY NOTICE:** This project is still under development and not stable yet - breaking changes may happen at any time and without notice

---

A Go client library enabling programs to perform CRUD operations on the [goharbor](https://github.com/goharbor/harbor) API.

This client library utilizes types generated by [go-swagger](https://github.com/go-swagger/go-swagger).

## Compatibility

This library includes separate clients supporting the **v1.10** & **v2.1** versions of Harbor. 

## Installation

Install the desired client library version using `go get`:

```shell script
# v1 Client
go get github.com/mittwald/goharbor-client/v3/apiv1
```

or

```shell script
# v2 Client
go get github.com/mittwald/goharbor-client/v3/apiv2
```

## Usage

Example usage of the client can be found in the [package examples](https://pkg.go.dev/github.com/mittwald/goharbor-client?tab=doc#pkg-examples).

## Documentation
For more specific documentation, please refer to the [godoc](https://pkg.go.dev/github.com/mittwald/goharbor-client) of this library.
