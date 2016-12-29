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
* [x] [Appdash](https://github.com/sourcegraph/appdash)


## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquet-0.2.0_netbsd_arm) ]

## Usage

Launch Zipkin with Docker, and open a browser on 9411:

    $ docker run -d -p 9411:9411 openzipkin/zipkin

or Appdash with Docker (open a browser on 7700):

    $ docker run -d -p 7700:7700 -p 7701:7701 solher/appdash

Launch the *trinquetd* server:

    $ ./trinquetd -config trinquet.toml -v 2 -logtostderr

Use the *trinquetctl* CLI to manage pelota informations :

    $ Trinquetctl is a CLI to use the Trinquet server

    Usage:
    trinquetctl [command]

    Available Commands:
        league      Print information about a league
        leagues     Print the available leagues
        version     Print the version number of Trinquetctl

    Use "trinquetctl [command] --help" for more information about a command.

You could explore the API using [Swagger](http://swagger.io/) UI :

    http://localhost:9090/swagger-ui/


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
