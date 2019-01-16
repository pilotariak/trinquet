# Trinquet

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/pilotariak%2Ftrinquet.svg)](https://badge.fury.io/gh/pilotariak%2Ftrinquet)

* Master : [![Circle CI](https://circleci.com/gh/pilotariak/trinquet/tree/master.svg?style=svg)](https://circleci.com/gh/pilotariak/trinquet/tree/master)
* Develop : [![Circle CI](https://circleci.com/gh/pilotariak/trinquet/tree/develop.svg?style=svg)](https://circleci.com/gh/pilotariak/trinquet/tree/develop)


*trinquet* uses [gRPC](http://www.grpc.io/) for its messaging protocol. The *trinquet* project includes a gRPC-based Go client and a command line utility, *trinquetctl*, for communicating with the *trinquetd* server.

For languages with no gRPC support, *trinquet* provides a [JSON](http://www.json.org/) grpc-gateway. This gateway serves a RESTful proxy that translates HTTP/JSON requests into gRPC messages.

It exports metrics for [Prometheus](https://prometheus.io/)

Distributed tracing is available using [OpenTracing](http://opentracing.io/). Supported tracers are:
* [x] [Zipkin](https://github.com/openzipkin)
* [x] [Jaeger](https://github.com/uber/jaeger)


## Installation

You can download the binaries :

### Trinquetd

* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_netbsd_arm) ]

### Trinquetctl

* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_netbsd_arm) ]

### Trinquetadm

* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetadm-0.2.0_netbsd_arm) ]

## Usage


## Development

### Kubernetes / Minikube

* Install Kubernetes tools (`minikube` and `kubectl` tools):

        $ make init

* Create the Kubernetes development cluster with minikube:

        $ make minikube-start
        $ make minikube-status
        $ make minikube-dashboard

### Database

* Install tools (`migrate` and `schemacrawler`):

        $ make init


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
