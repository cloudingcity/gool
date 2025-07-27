# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gool is a Go-based CLI toolkit inspired by Boop for common developer tasks like string transformations, date conversions, JWT decoding, URL encoding/decoding, and more. It operates in dual modes: direct command execution and interactive shell mode. The project is designed as both a CLI tool and a reusable Go library.

## Development Commands

### Building and Testing
- `go build` - Build the binary
- `go run main.go <command> <input>` - Run a specific command directly
- `go run main.go` - Start default shell mode (no args defaults to shell)
- `make test` - Run all tests with race detection and atomic coverage
- `make lint` - Run golangci-lint (auto-installs v1.30.0 if not present)

### Installation Methods
- **Homebrew**: `brew install cloudingcity/tap/gool`
- **Docker**: `docker run --rm -it ghost0436/gool`
- **Go install**: Standard go build/install workflow

## Architecture

### Layered Architecture Pattern
```
main.go → cmd/ (Cobra commands) → pkg/ (business logic)
           ↓
    internal/shell/ (interactive mode)
```

### Dual Operation Modes
1. **Direct CLI**: `gool jwt-decode <token>` - Single command execution
2. **Interactive Shell**: `gool` or `gool shell` - REPL with command switching

### Command Implementation Pattern
Every utility follows this consistent structure:
- **Command Definition**: `cmd/<utility>.go` with Cobra command setup
- **Business Logic**: `pkg/<category>/` with pure functions
- **Testing**: `*_test.go` with table-driven tests using testify

### Package Organization Philosophy
- **`pkg/`** - Pure business logic, importable as library, no CLI dependencies
- **`cmd/`** - Thin Cobra wrappers with input validation and output formatting
- **`internal/shell/`** - Shell-specific logic (readline, command routing, history)

### Key Dependencies & Their Roles
- **`github.com/spf13/cobra v1.9.1`** - CLI framework with command routing and help
- **`github.com/chzyer/readline v1.5.1`** - Interactive shell with history/autocomplete
- **`github.com/araddon/dateparse`** - Flexible date parsing (handles many formats)
- **`github.com/tidwall/pretty v1.2.1`** - JSON formatting with colorization
- **`github.com/fatih/color v1.18.0`** - Terminal color output

### Command Registration System
All commands are centrally registered via the `commands()` function in `cmd/root.go`:
```go
func commands() []*cobra.Command {
    return []*cobra.Command{
        timestampToDateCmd,  // alias: "t2d"
        dateToTimestampCmd,  // alias: "d2t"
        // ... all other commands
    }
}
```

Commands support aliases and follow `DisableFlagsInUseLine: true` convention.

### Shell Mode Architecture
The shell mode (`internal/shell/`) provides:
- **Command Switching**: `\s <command>` changes active script context
- **Direct Execution**: `\r <command> <args>` runs without switching
- **Built-in Commands**: `\h` (help), `\l` (list), `\c` (clear), `\q` (quit)
- **History Persistence**: Saved to temp directory
- **Autocomplete**: Commands and script names
- **Command Reuse**: Shell registers all CLI commands automatically

### Testing Strategy
- **Table-driven tests**: Consistent pattern across all packages
- **Testify assertions**: `assert.Equal`, `assert.NoError` for readable tests
- **Race detection**: All tests run with `-race` flag in CI
- **Coverage tracking**: Atomic mode with Codecov integration
- **Integration tests**: Shell package tests full interaction flows

### Build & Deployment
- **Multi-platform**: GoReleaser builds for Linux, Windows, macOS
- **Distribution**: GitHub releases, Homebrew tap, Docker images
- **CI/CD**: GitHub Actions with test, lint, and deploy workflows
- **Quality gates**: golangci-lint, codecov, go report card

### Extension Patterns
To add a new utility:
1. Create business logic function in appropriate `pkg/<category>/`
2. Add table-driven tests for the function
3. Create Cobra command in `cmd/<utility>.go`
4. Add command to `commands()` function in `cmd/root.go`
5. Shell mode automatically includes the new command

### Library Usage
The `pkg/` directory is designed for external consumption:
```go
import "github.com/cloudingcity/gool/pkg/cases"
cases.Snake("HelloWorld") // "hello_world"
```

### Version Management
Version and build date are injected at compile time via `main.go` and displayed in `cmd/version.go`.