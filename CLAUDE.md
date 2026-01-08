# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library providing generic utility functions for working with slices. The library uses Go generics extensively to provide type-safe operations on slices of any type.

## Building and Testing

### Build
```bash
go build -v ./...
```

### Run all tests
```bash
go test -v ./...
```

### Run tests for a specific function
```bash
go test -v -run TestFeatures/Filter
go test -v -run TestFeatures/Flatten
```

### Run a specific BDD scenario
The project uses Cucumber/Godog for BDD testing. Feature files are in the `features/` directory.
```bash
go test -v -run TestFeatures
```

### Linting
The project uses pre-commit hooks with golangci-lint:
```bash
golangci-lint run
```

### Format code
```bash
go fmt ./...
```

### Tidy dependencies
```bash
go mod tidy
```

## Code Architecture

### Testing Architecture
This codebase uses **Behavior-Driven Development (BDD)** with Cucumber/Godog as the primary testing approach, not traditional Go unit tests.

- **Feature files**: Located in `features/` directory, written in Gherkin syntax (e.g., `filter.feature`, `flatten.feature`)
- **Test orchestration**: `slices_test.go` contains the main test suite runner and shared step definitions
- **Function-specific tests**: Each function has its own `*_test.go` file (e.g., `filter_test.go`) that implements step definitions specific to that function
- **Context-based state**: Tests use Go's `context.Context` with custom keys (`firstArgKey`, `resultKey`) to pass data between step definitions
- **Test suite registration**: `InitializeScenario()` in `slices_test.go` registers all step definitions from individual test files

### File Organization
Each slice utility function follows a consistent pattern:
- `<function>.go` - Implementation (e.g., `filter.go`, `flatten.go`)
- `<function>_test.go` - Godog step definitions (e.g., `filter_test.go`)
- `features/<function>.feature` - BDD scenarios in Gherkin format

### Adding New Functions
When adding a new slice utility function:

1. Create `<function>.go` with the implementation
2. Create `<function>_test.go` with Godog step definitions:
   ```go
   func initializeScenarioFor<Function>(ctx *godog.ScenarioContext) {
       ctx.Step(`^step definition regex$`, stepFunction)
   }
   ```
3. Create `features/<function>.feature` with BDD scenarios
4. Register the initializer in `slices_test.go` in the `InitializeScenario()` function:
   ```go
   initializeScenarioFor<Function>(ctx)
   ```

### Shared Test Utilities
`slices_test.go` provides reusable step definitions and helper functions:
- `tableToIntegerSlice()` - Convert Godog tables to `[]int`
- `tableToStringSlice()` - Convert Godog tables to `[]string`
- `anIntegerSliceWithElements()` - Common Given step for integer slices
- `theResultShouldBeAnIntegerSlice()` - Common Then step for validating results

Use these shared utilities when writing new test scenarios to maintain consistency.

## Pre-commit Hooks

The repository uses pre-commit hooks configured in `.pre-commit-config.yaml`:
- `go-fmt` - Format Go code
- `golangci-lint` - Lint Go code
- `go-unit-tests` - Run tests (actually runs BDD tests via `go test`)
- `go-mod-tidy` - Ensure go.mod is tidy
- `commitlint` - Enforce conventional commit messages
- `reformat-gherkin` - Format .feature files

## Conventional Commits

This project uses conventional commits enforced by commitlint. Commit messages must follow the format:
```
type(scope): description

Examples:
feat: add new Map function
fix: correct nil handling in FilterNil
test: add scenarios for edge cases
docs: update README with new examples
```

## Go Version

This project requires **Go 1.24.0** or later (specified in `go.mod`).

## Branch Strategy

- Main development branch: `develop`
- Production branch: `main`
- PRs should target `develop`