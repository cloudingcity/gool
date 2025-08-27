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
- Direct command execution without mode switching (e.g., `/run uuid 5`)
- History management with temp directory storage
- Built-in commands: `/help`, `/run`, `/clear`, `/quit`

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
- **Date/Time**: tsconv (smart timestamp/date converter), mstsconv (millisecond version), ts2date, date2ts  
- **Encoding**: base64-encode/decode, url-encode/decode, jwt-decode
- **Text Processing**: case conversions (camel, kebab, snake, etc.), count, text-escape, text-unescape
- **Utilities**: md5, uuid, format-json, url-to-json

### tsconv Command
The `tsconv` command intelligently converts between timestamps and dates:
- **No input**: Shows current time as both timestamp and date
- **Timestamp input**: Converts Unix timestamp (seconds) to RFC3339 date format
- **Date input**: Converts date string to Unix timestamp (seconds)
- **Examples**:
  ```bash
  gool tsconv                    # Show current time
  gool tsconv 1609459200         # Convert timestamp to date
  gool tsconv "2021-01-01"       # Convert date to timestamp
  ```

### mstsconv Command
The `mstsconv` command works with millisecond timestamps:
- **No input**: Shows current time as both millisecond timestamp and date
- **Millisecond timestamp input**: Converts to RFC3339 date format
- **Date input**: Converts date string to millisecond timestamp
- **Examples**:
  ```bash
  gool mstsconv                  # Show current time
  gool mstsconv 1609459200000    # Convert MS timestamp to date
  gool mstsconv "2021-01-01"     # Convert date to MS timestamp
  ```

### /run Command (Interactive Shell Only)
The `/run` command allows direct execution of any command without switching modes in the interactive shell:
- **Syntax**: `/run <command> [args...]`
- **Auto-completion**: Tab completion works for both command names and arguments
- **No mode switching**: Commands execute immediately and return to main prompt
- **Help**: Use `/run` without arguments to see available commands
- **Examples**:
  ```bash
  /run uuid 5                    # Generate 5 UUIDs
  /run md5 hello                 # Compute MD5 hash of "hello"
  /run tsconv 1609459200         # Convert timestamp to date
  /run camel-case hello_world    # Convert to camelCase
  /run base64-encode "text"      # Base64 encode text
  ```

## Key Dependencies

- **Cobra**: CLI framework for command structure
- **Readline**: Interactive shell with history and auto-completion
- **Color**: Terminal color output
- **Dateparse**: Flexible date parsing
- **Pretty**: JSON formatting

## Maintenance Tasks

- **Update README.md or CLAUDE.md if necessary**

## Development Reminders

- Ensure make lint is not failed