# branch-cleaner

[![Release](https://img.shields.io/github/release/knipferrc/branch-cleaner.svg?style=flat-square)](https://github.com/knipferrc/branch-cleaner/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/knipferrc/branch-cleaner?style=flat-square)](https://goreportcard.com/report/github.com/knipferrc/branch-cleaner)
[![Godoc](https://godoc.org/github.com/knipferrc/branch-cleaner?status.svg&style=flat-square)](http://godoc.org/github.com/knipferrc/branch-cleaner)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

![Screenshot](/assets/screenshot.png)

## About The Project

A TUI to cleanup local git branches

### Built With

- [Go](https://golang.org/)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
- [Viper](https://github.com/spf13/viper)
- [Cobra](https://github.com/spf13/cobra)

## Installation

### Curl

```sh
curl -sfL https://raw.githubusercontent.com/knipferrc/branch-cleaner/main/install.sh | sh
```

### Go

```
go install github.com/knipferrc/branch-cleaner@latest
```

## Features

- Easily delete branches one at a time or in bulk
- Ability to filter by branch name
- Protect branches from being deleted

## Configuration

- A config file will be generated at `~/.branch-cleaner.yml` when you first run `branch-cleaner`

```yml
settings:
  enable_logging: false
  protected_branches:
    - main
    - master
    - develop
    - dev
    - prod
```
