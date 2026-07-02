# src

This directory contains the source code for the `common` package.

## Contents

| Directory | Language | Description |
|-----------|----------|-------------|
| `go/`     | Go       | Shared types, utilities, and configuration used across History Tracers projects |

The `go/` directory is a Go module (`github.com/historytracers/common`) exposing the `common` package with data structures, configuration types, timestamps, and indexing helpers for class and family records.

## Building

```bash
# Regenerate autotools files (if configure.ac or Makefile.am changed)
./bootstrap

# Configure and build
./configure
make
```
