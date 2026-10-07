# publisher

Local publisher for the smartphone content stored in `src/smartphone/??-??/`.

It applies the same transformations used by the main History Tracers
publisher (default smileys, `source_menu` page URLs and, when a file is
modified in Git, its `last_update` timestamp) and writes the minified JSON to
the content directory.

## Build

The helper script at the repository root bootstraps the build system (when
needed), configures, builds the publisher, and then rewrites and verifies the
smartphone content:

```bash
./build-publisher.sh
```

It is equivalent to, from the repository root:

```bash
./bootstrap      # first time only, generates the build system
./configure
make publisher   # produces build/historytracers-publisher(.exe)
./build/historytracers-publisher -minify
./build/historytracers-publisher -validate
```

## Usage

Run the binary from the repository root:

```bash
# Rewrite and minify every smartphone file into build/www/lang/<lang>/smartphone/
./build/historytracers-publisher -minify

# Validate smartphone JSON files without producing output
./build/historytracers-publisher -validate

# Show the directories in use
./build/historytracers-publisher -compilation
```

### Options

| Flag | Description |
| --- | --- |
| `-minify` | Rewrite and minify all smartphone JSON files into the content directory. |
| `-validate` | Validate smartphone JSON files without producing output; exits non-zero when a file is invalid. |
| `-verbose` | Print information messages during processing. |
| `-compilation` | Show the directories used by this build. |
| `-src` | Directory containing the source files (default: `.`). |
| `-www` | Directory for the generated content (default: `build/www/`). |
| `-logfile` | Redirect all output to the given file. |
| `-db` | Path to the academic sources SQLite database (optional). |

When the optional sources database (`lang/sources/history_tracers.db` by
default) is present, citation data referenced by the smartphone files is
loaded while publishing. The publisher works without it.
