# publisher

Local publisher for the smartphone content stored in `src/smartphone/??-??/`.

It applies the same transformations used by the main History Tracers publisher
(default smileys, `source_menu` page URLs and, when a file is modified in Git,
its `last_update` timestamp), writes the minified JSON to the content directory
and generates the text-to-speech input files.

## Build

The helper script at the repository root bootstraps the build system (when
needed), configures, builds the publisher, and then rewrites, generates the
audio and verifies the smartphone content:

```bash
./build-publisher.sh
```

It is equivalent to, from the repository root:

```bash
./bootstrap      # first time only, generates the build system
./configure
make publisher   # produces build/historytracers-publisher(.exe)
./build/historytracers-publisher -minify
./build/historytracers-publisher -audio
./build/historytracers-publisher -validate
```

## Usage

Run the binary from the repository root:

```bash
# Rewrite and minify every smartphone file into build/www/lang/<lang>/smartphone/
./build/historytracers-publisher -minify

# Generate text-to-speech input files into audio/
./build/historytracers-publisher -audio

# Validate smartphone JSON files without producing output
./build/historytracers-publisher -validate

# Show the directories in use
./build/historytracers-publisher -compilation
```

### Options

| Flag | Description |
| --- | --- |
| `-minify` | Rewrite and minify all smartphone JSON files into the content directory. |
| `-audio` | Generate text-to-speech input files for every smartphone screen into the audio directory. |
| `-validate` | Validate smartphone JSON files without producing output; exits non-zero when a file is invalid. |
| `-verbose` | Print information messages during processing. |
| `-compilation` | Show the directories used by this build. |
| `-src` | Directory containing the source files (default: `.`). |
| `-www` | Directory for the generated content (default: `build/www/`). |
| `-audiodir` | Directory for the generated audio text files (default: `audio/`). |
| `-logfile` | Redirect all output to the given file. |
| `-db` | Path to the academic sources SQLite database (optional). |

When the optional sources database (`lang/sources/history_tracers.db` by
default) is present, citation data referenced by the smartphone files is
loaded while publishing. The publisher works without it.

## Audio files

`-audio` writes one plain-text file per screen, named
`audio/<lesson-uuid>_<screen-uuid>_<lang>.txt`. The files are intended as input
for a text-to-speech engine such as [Piper](https://github.com/rhasspy/piper).
HTML and markdown markup, emoji and social-media URLs are simplified so the
text reads naturally. The output directory can be changed with `-audiodir`.

## Reusing the package

The rewrite, minify and audio logic lives in the importable package
`historytracers-publisher/smartphone`, so other History Tracers projects can
call it directly:

```go
import "historytracers-publisher/smartphone"

cfg := smartphone.Config{
    SrcPath:     "/path/to/common",
    ContentPath: "/path/to/output",
    AudioPath:   "/path/to/audio",
}

if err := smartphone.HTMinifyAllFiles(cfg); err != nil {
    // handle error
}

if err := smartphone.HTGenerateAudio(cfg); err != nil {
    // handle error
}

if invalid := smartphone.HTValidateSMGameFormats(cfg); invalid > 0 {
    // handle invalid files
}
```

Consumers reference this module with a `replace` directive (the same pattern
used for `github.com/historytracers/common`):

```
require historytracers-publisher v0.0.0

replace historytracers-publisher => ../common/src/publisher
```

The command in this directory is only a thin wrapper around the package.
