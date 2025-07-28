# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Gool is a Go-based CLI toolkit that provides utilities for common programming tasks (inspired by Boop). It operates in two modes:
1. **Interactive Shell Mode**: Start with `gool` for a REPL-like interface
2. **Direct CLI Mode**: Run commands directly like `gool jwt-decode <token>`

## Architecture

### Core Components

- **`main.go`**: Entry point that sets version and executes the root command
- **`cmd/`**: Contains all command implementations using Cobra CLI framework
  - `root.go`: Root command setup and command registration
  - Individual command files (e.g., `jwt.go`, `case.go`, etc.)
- **`pkg/`**: Reusable library packages that can be imported by other Go projects
  - Each utility has its own package (cases, date, jwt, random, timestamp, url)
- **`internal/shell/`**: Interactive shell implementation using readline
  - `shell.go`: Core shell logic with command switching and auto-completion
  - `reader.go`: Readline configuration and history management

### Command Architecture

Commands are implemented using the Cobra CLI framework. Each command:
1. Is defined in `cmd/` directory 
2. Is registered in the `commands()` function in `cmd/root.go`
3. Can be used both in interactive shell mode and direct CLI mode
4. Has corresponding library functions in `pkg/` for programmatic use

### Shell System

The interactive shell (`internal/shell/`) provides:
- Command auto-completion with `/` prefix
- Command switching (e.g., `/jwt-decode` switches to JWT decode mode)
- History management with temp directory storage
- Built-in commands: `/help`, `/clear`, `/quit`

## Development Commands

### Build & Run
```bash
go build -o gool .        # Build binary
./gool                    # Run interactive mode
./gool <command> <args>   # Run direct mode
```

### Testing
```bash
make test                 # Run all tests with coverage
go test ./...             # Standard test run
go test -v ./pkg/cases/   # Test specific package
```

### Linting
```bash
make lint                 # Run golangci-lint (installs if needed)
```

### Individual Package Testing
```bash
go test -v ./pkg/jwt/      # Test JWT package
go test -v ./pkg/cases/    # Test case conversion package
go test -v ./internal/shell/  # Test shell functionality
```

## Available Commands

The tool provides utilities for:
- **Date/Time**: timestamp-to-date, date-to-timestamp  
- **Encoding**: base64-encode/decode, url-encode/decode, jwt-decode
- **Text Processing**: case conversions (camel, kebab, snake, etc.), count
- **Utilities**: md5, rand-string, format-json, url-to-json

## Key Dependencies

- **Cobra**: CLI framework for command structure
- **Readline**: Interactive shell with history and auto-completion
- **Color**: Terminal color output
- **Dateparse**: Flexible date parsing
- **Pretty**: JSON formatting