#!/bin/sh
# Build the local smartphone publisher and regenerate the published content.
#
# The script generates the build system when needed, runs configure, builds
# only the publisher (build/historytracers-publisher[.exe]), then rewrites and
# verifies every file under src/smartphone/??-??/. Any extra arguments are
# forwarded to configure, for example:
#
#   ./build-publisher.sh --with-go-compiler=/path/to/go

set -e

# Always run from the repository root, regardless of the caller's directory.
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
cd "$SCRIPT_DIR"

# Detect the binary extension used by the host platform.
case "$(uname -s)" in
    CYGWIN*|MINGW*|MSYS*)
        EXEEXT=".exe"
        ;;
    *)
        EXEEXT=""
        ;;
esac

# Generate the autoconf/automake files on a fresh checkout.
if [ ! -x ./configure ]; then
    echo "=== Bootstrapping the build system ==="
    ./bootstrap
fi

echo "=== Configuring ==="
./configure "$@"

echo "=== Building publisher ==="
make publisher

PUBLISHER_BIN="$SCRIPT_DIR/build/historytracers-publisher$EXEEXT"

echo "=== Rewriting smartphone content ==="
"$PUBLISHER_BIN" -minify

echo "=== Verifying smartphone content ==="
"$PUBLISHER_BIN" -validate

echo ""
echo "Publisher built and smartphone content regenerated: $PUBLISHER_BIN"
