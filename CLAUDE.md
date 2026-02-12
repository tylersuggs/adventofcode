# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains solutions to Advent of Code challenges for years 2020, 2021, and 2025. Each day's puzzle is self-contained in its own directory with a simple CLI solution.

## Project Structure

- Days are organized by year in directories (e.g., `2020/01`, `2021/01`, `2025/01`)
- Each day directory contains:
  - `main.go` - the primary solution code
  - `input.txt` or `input.go` - the puzzle input data (input.go contains the raw input as a Go string literal)
  - `go.mod` - Go module definition for that day
  - Some days have a `bin/` directory (build output, ignored by git)

## Development Commands

To build and run a specific day's solution:

```bash
cd 2020/01
go run main.go
```

Or to compile to a binary first:

```bash
cd 2020/01
go build -o bin/solution main.go
./bin/solution
```

Note: Each day is its own Go module, so you need to `cd` into the day's directory before running Go commands.

## Code Style

Solutions are simple, self-contained programs that:
1. Read input from `input.txt` or use hardcoded input from `input.go`
2. Implement the day's puzzle logic
3. Print answers directly to stdout
4. Follow standard Go conventions

No complex testing frameworks or linting configuration is used. If tests are needed, standard Go testing (`*_test.go`) can be added within the same package.