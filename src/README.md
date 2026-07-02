# src

This directory contains the source code for the `common` package.

## Contents

| Directory     | Language        | Description |
|---------------|-----------------|-------------|
| `android/`    | Java (Android)  | Android library with data models, timestamp utilities, and creator helpers for History Tracers apps |
| `go/`         | Go              | Shared types, utilities, and configuration used across History Tracers projects |

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
