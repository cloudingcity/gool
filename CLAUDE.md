# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gool is a Go-based CLI toolkit for common developer tasks like string transformations, date conversions, JWT decoding, URL encoding/decoding, and more. It can run individual commands or operate in an interactive shell mode.

## Development Commands

### Building and Testing
- `go build` - Build the binary
- `go run main.go <command>` - Run a specific command
- `make test` - Run all tests with race detection and coverage
- `make lint` - Run golangci-lint (installs if not present)

### Running the Application
- `go run main.go <command> <input>` - Execute a specific utility command
- `go run main.go shell` - Start interactive shell mode
- `go run main.go --help` - Show all available commands

## Architecture

### Command Structure
- **`main.go`** - Entry point that sets version/date and calls cmd.Execute()
- **`cmd/root.go`** - Cobra root command setup and command registration
- **`cmd/*.go`** - Individual command implementations (jwt, base64, case conversions, etc.)
- **`internal/shell/`** - Interactive shell implementation with readline support

### Package Organization
- **`pkg/`** - Reusable utility packages that can be imported as a library:
  - `cases/` - String case conversions (camel, snake, kebab, etc.)
  - `date/` - Date parsing and conversion utilities
  - `jwt/` - JWT token decoding
  - `random/` - Random string generation
  - `timestamp/` - Unix timestamp conversions
  - `url/` - URL encoding/decoding and parsing

### Key Dependencies
- `github.com/spf13/cobra` - CLI framework
- `github.com/chzyer/readline` - Interactive shell with history/autocomplete
- `github.com/fatih/color` - Terminal color output
- `github.com/araddon/dateparse` - Flexible date parsing
- `github.com/tidwall/pretty` - JSON formatting

## Shell Mode
The interactive shell (`gool shell`) provides:
- Command history and autocomplete
- Script switching with `\s <command>`
- Direct execution with `\r <command> <args>`
- Built-in help with `\h` and script listing with `\l`

## Testing
All packages have corresponding test files using standard Go testing and testify assertions. Tests include both unit tests for individual functions and integration tests for command execution.