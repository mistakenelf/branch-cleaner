<p align="center">
  <h1 align="center">branch-cleaner</h1>
  <p align="center">
    <a href="https://github.com/knipferrc/branch-cleaner/releases"><img src="https://img.shields.io/github/v/release/knipferrc/branch-cleaner" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/knipferrc/branch-cleaner?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/knipferrc/branch-cleaner/actions"><img src="https://img.shields.io/github/workflow/status/knipferrc/branch-cleaner/Release" alt="Build Status"></a>
  </p>
</p>

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

- TUI to view the current git directories branches and ability to delete them easily

## Configuration

- A config file will be generated at `~/.branch-cleaner.yml` when you first run `branch-cleaner`

```yml
settings:
  enable_logging: false
  enable_mousewheel: true
```
