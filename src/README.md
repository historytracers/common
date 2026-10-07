# src

This directory contains the source code for the `common` package.

## Contents

| Directory     | Language        | Description |
|---------------|-----------------|-------------|
| `android/`    | Java (Android)  | Android library with data models, timestamp utilities, and creator helpers for History Tracers apps |
| `go/`         | Go              | Shared types, utilities, and configuration used across History Tracers projects |
| `smartphone/` | JSON            | Smartphone lesson content, one directory per language |
| `publisher/`  | Go              | Local publisher that rewrites and minifies the smartphone JSON files |

### android/

Android library module exposing Java POJOs (with Gson serialization) mirroring the Go data structures. Includes `HTConfigBase`, all content model classes (`Family`, `ClassIdx`, `ClassTemplateFile`, `HTSourceFile`, etc.), `TimestampUtils`, and `CreatorUtils`.

Build with Gradle:

```bash
./gradlew :common-android:assembleRelease
```

### go/

Go module (`github.com/historytracers/common`) exposing the `common` package with data structures, configuration types, timestamps, and indexing helpers for class and family records.

Build with autotools:

```bash
./bootstrap
./configure
make
```

### smartphone/

Smartphone lesson content, one directory per language (`en-US/`, `es-ES/`,
`pt-BR/`) with one JSON file per lesson.

### publisher/

Go command that rewrites (normalizes) and minifies the smartphone JSON files.

Build it with the helper script from the repository root:

```bash
./build-publisher.sh
```

Rewrite (normalize and minify) every JSON file:

```bash
./build/historytracers-publisher -minify
```

The generated files are written to `build/www/lang/<lang>/smartphone/`; the
source files are not modified. Verify the JSON files with:

```bash
./build/historytracers-publisher -validate
```

See [publisher/README.md](publisher/README.md) for details.
