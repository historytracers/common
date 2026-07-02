#!/bin/bash

set -eo pipefail

# Detect platform
case "$(uname -s)" in
    Linux)
        PLATFORM="linux"
        ;;
    CYGWIN*|MINGW*|MSYS*)
        PLATFORM="msys2"
        if [ -z "$GOROOT" ]; then
            for cand in "/c/Program Files/Go" "/c/Go" "/mingw64/lib/go"; do
                if [ -x "$cand/bin/go.exe" ] || [ -x "$cand/bin/go" ]; then
                    export GOROOT="$cand"
                    break
                fi
            done
        fi
        ;;
    *)
        echo "Unknown platform: $(uname -s)"
        exit 1
        ;;
esac

compile() {
    echo "=== Platform: $PLATFORM ==="

    autoreconf -f -i
    echo "=== autoreconf done ==="

    ./configure
    echo "=== configure done ==="

    make clean
    make all
    echo "=== build done ==="

    make test
    echo "=== tests done ==="
}

compile
