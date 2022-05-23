# ![icon](./docs/images/icon.png) Sherlock
[![Go Report Card](https://goreportcard.com/badge/github.com/jwillker/sherlock)](https://goreportcard.com/report/github.com/jwillker/sherlock)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jwillker/sherlock/main/LICENSE)
[![Golang](https://img.shields.io/badge/Go-1.18-blue.svg)](https://golang.org)
[![build](https://github.com/jwillker/sherlock/actions/workflows/build.yml/badge.svg)](https://github.com/jwillker/sherlock/actions/workflows/build.yml)
[![golangci-lint](https://github.com/jwillker/sherlock/actions/workflows/lint.yml/badge.svg)](https://github.com/jwillker/sherlock/actions/workflows/lint.yml)

A tool to search and index Golang projects in [Sally](https://github.com/uber-go/sally) vanity config

## Install

```
go install github.com/jwillker/sherlock@latest 
```

or download the [binary](https://github.com/jwillker/sherlock/releases) specific for your OS.

## How to use

:warning: For now it only supports gitlab

```
Usage:
  sherlock investigate [flags]

Flags:
  -u, --base-url string     GitLab api url. Default https://gitlab.com/api/v4
  -d, --godoc string        Godoc URL
  -g, --group-id int        GitLab group id
  -h, --help                help for investigate
  -o, --output string       Vanity packages config. Default: packages.yaml (default "packages.yaml")
  -v, --vanity-url string   Sally vanity URL
```


## Example

Running:

`sherlock investigate -u https://gitlab.my-company.com/api/v4 -g 123 -d godoc.my-company.com/pkg -v go.my-company.com`

Result:

`cat packages.yaml`

```yaml
godoc:
    host: godoc.my-company.com/pkg
url: go.my-company.com
packages:
    logs:
        repo: https://gitlab.my-company.com/libs/golang/logs.git
    sqs:
        repo: https://gitlab.my-company.com/libs/golang/sqs.git
```

So you can host this using [Sally](https://github.com/uber-go/sally) vanity by running:

`sally -yml packages.yaml -port 8080`

## Developing
Use make help to see what to run, some options:

```table
help                           Show Help
build                          Build sherlock
dep                            Get the dependencies
lint                           Run linter
race                           Run data race detector
msan                           Run memory sanitizer
test                           Run tests
```
