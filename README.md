# ![icon](./docs/images/icon.png) Sherlock
[![Go Report Card](https://goreportcard.com/badge/github.com/jwillker/sherlock)](https://goreportcard.com/report/github.com/jwillker/sherlock)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jwillker/sherlock/main/LICENSE)
[![Golang](https://img.shields.io/badge/Go-1.18-blue.svg)](https://golang.org)

A tool to search and index Golang projetcs in Sally vanity config

## How to use

```
Usage:
  sherlock investigate [flags]

Flags:
  -u, --base-url string   GitLab api url. Default https://gitlab.com/api/v4
  -g, --group-id int      GitLab group id
  -h, --help              help for investigate
  -o, --output string     Vanity packages config. Default: packages.yaml (default "packages.yaml")
```

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
