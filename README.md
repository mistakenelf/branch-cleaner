# branch-cleaner

[![Release](https://img.shields.io/github/release/mistakenelf/branch-cleaner.svg?style=flat-square)](https://github.com/mistakenelf/branch-cleaner/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/mistakenelf/branch-cleaner?style=flat-square)](https://goreportcard.com/report/github.com/mistakenelf/branch-cleaner)
[![Godoc](https://godoc.org/github.com/mistakenelf/branch-cleaner?status.svg&style=flat-square)](http://godoc.org/github.com/mistakenelf/branch-cleaner)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

![Screenshot](/assets/screenshot.png)

## About The Project

A TUI to cleanup local git branches

### Built With

- [Go](https://golang.org/)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
- [Cobra](https://github.com/spf13/cobra)

## Installation

### Curl

```sh
curl -sfL https://raw.githubusercontent.com/mistakenelf/branch-cleaner/main/install.sh | sh
```

### Go

```
go install github.com/mistakenelf/branch-cleaner@latest
```

## Features

- Easily delete branches one at a time or in bulk
- Ability to filter by branch name
- Protect branches from being deleted

## Usage

- `branch-cleaner` will start the application in the current directory
- `branch-cleaner -h` will show the help menu
- `branch-cleaner -v` will show the current version
- `branch-cleaner update` will update branch cleaner to the latest version

Too much to type? Create an alias:

```
alias bc="branch-cleaner"
```

## Configuration

Configuration is handled through environment variables.

```sh
PROTECTED_BRANCHES="main:master:develop:dev:prod" branch-cleaner
```
