#!/bin/bash

set -eo pipefail

# Detect platform
case "$(uname -s)" in
    Linux)
        PLATFORM="linux"
        ;;
    CYGWIN*|MINGW*|MSYS*)
        PLATFORM="msys2"

        # Prefer the UCRT64 environment when this MSYS2 installation provides it.
        # UCRT64 is the recommended toolchain on modern MSYS2; switching to it
        # early makes configure/make pick the UCRT64 Go compiler and MinGW-w64
        # C compiler instead of the default MSYS2 ones. When /ucrt64 is missing
        # (or has no toolchain), fall back to the default MSYS2 environment.
        MSYS2_ENV="${MSYSTEM:-MSYS}"
        if [ "$MSYS2_ENV" != "UCRT64" ] && [ -d /ucrt64 ]; then
            if [ -x /ucrt64/bin/x86_64-w64-mingw32-gcc ] \
                || [ -x /ucrt64/bin/gcc ] \
                || [ -x /ucrt64/bin/go.exe ] \
                || [ -x /ucrt64/bin/go ]; then
                export MSYSTEM=UCRT64
                export MINGW_PREFIX=/ucrt64
                export MINGW_CHOST=x86_64-w64-mingw32
                export MINGW_PACKAGE_PREFIX=mingw-w64-ucrt-x86_64
                case ":$PATH:" in
                    *":/ucrt64/bin:"*) ;;
                    *) export PATH="/ucrt64/bin:$PATH" ;;
                esac
                export PKG_CONFIG_PATH="/ucrt64/lib/pkgconfig:/ucrt64/share/pkgconfig${PKG_CONFIG_PATH:+:$PKG_CONFIG_PATH}"
                export ACLOCAL_PATH="/ucrt64/share/aclocal${ACLOCAL_PATH:+:$ACLOCAL_PATH}"
                MSYS2_ENV="UCRT64"
            else
                echo "WARNING: /ucrt64 exists but no UCRT64 toolchain found; using default MSYS2 environment ($MSYS2_ENV)."
            fi
        fi

        # Set GOROOT if not set and Go is in a known location
        if [ -z "$GOROOT" ]; then
            for cand in "/ucrt64/lib/go" "/c/Program Files/Go" "/c/Go" "/mingw64/lib/go"; do
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
    echo "=== Platform: $PLATFORM${MSYS2_ENV:+ ($MSYS2_ENV)} ==="

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
