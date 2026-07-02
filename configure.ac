# Initialize Autoconf
AC_INIT([common], [1.0.0], [historytracers@gmail.com])

# Set package name and version
AC_CONFIG_SRCDIR([src/go/config.go])

# Initialize Automake
AM_INIT_AUTOMAKE([foreign subdir-objects])

# Check for required programs
AC_PROG_INSTALL
AC_PROG_MKDIR_P
AC_PROG_MAKE_SET

# Check for Go compiler
# Go compiler selection options
AC_ARG_WITH([go-compiler],
    [AS_HELP_STRING([--with-go-compiler=COMPILER],
        [Specify Go compiler to use (go, gccgo, or full path) @<:@auto@:>@])],
    [go_compiler="$withval"],
    [go_compiler="auto"])

AC_ARG_WITH([go-path],
    [AS_HELP_STRING([--with-go-path=PATH],
        [Specify directory containing Go compiler])],
    [go_path="$withval"],
    [go_path=""])

# Set up PATH if custom Go path specified
if test "x$go_path" != "x"; then
    save_PATH="$PATH"
    PATH="$go_path:$PATH"
    export PATH
fi

# Check for Go compiler
AC_MSG_CHECKING([for Go compiler])
GO_BIN_DIR=""
if test "x$go_compiler" != "xauto" -a "x$go_compiler" != "x"; then
    if test -x "$go_compiler"; then
        GO_COMPILER="$go_compiler"
        AC_MSG_RESULT([$GO_COMPILER (user specified)])
    else
        AC_CHECK_PROG([HAVE_USER_GO], [$go_compiler], [yes], [no])
        if test "x$HAVE_USER_GO" = "xyes"; then
            GO_COMPILER="$go_compiler"
            AC_MSG_RESULT([$GO_COMPILER (user specified)])
        else
            AC_MSG_ERROR([User-specified Go compiler '$go_compiler' not found])
        fi
    fi
else
    AC_CHECK_PROG([HAVE_GO], [go], [yes], [no])
    if test "x$HAVE_GO" = "xyes"; then
        GO_COMPILER="go"
        AC_MSG_RESULT([$GO_COMPILER (auto-detected)])
    else
        AC_CHECK_PROG([HAVE_GCCGO], [gccgo], [yes], [no])
        if test "x$HAVE_GCCGO" = "xyes"; then
            GO_COMPILER="gccgo"
            AC_MSG_RESULT([$GO_COMPILER (auto-detected)])
        else
            # Fallback: try common Windows Go installation paths (MSYS/MinGW)
            if test -x "/c/Program Files/Go/bin/go.exe"; then
                GO_BIN_DIR="/c/Program Files/Go/bin"
                PATH="$GO_BIN_DIR:$PATH"
                GO_COMPILER="go"
                AC_MSG_RESULT([$GO_COMPILER ($GO_BIN_DIR)])
            elif test -x "/c/Go/bin/go.exe"; then
                GO_BIN_DIR="/c/Go/bin"
                PATH="$GO_BIN_DIR:$PATH"
                GO_COMPILER="go"
                AC_MSG_RESULT([$GO_COMPILER ($GO_BIN_DIR)])
            else
                GO_RAW=`cmd.exe /c "where go" 2>/dev/null | head -1 | tr -d '\r'`
                if test -n "$GO_RAW"; then
                    GO_MSYS=`cygpath -u "$GO_RAW" 2>/dev/null || echo "$GO_RAW" | sed 's/\\/\//g; s/^C:/\/c/'`
                    GO_BIN_DIR=`dirname "$GO_MSYS" 2>/dev/null`
                    if test -d "$GO_BIN_DIR" && test -x "$GO_MSYS"; then
                        PATH="$GO_BIN_DIR:$PATH"
                        GO_COMPILER="go"
                        AC_MSG_RESULT([$GO_COMPILER (from Windows PATH)])
                    else
                        AC_MSG_RESULT([not usable ($GO_MSYS)])
                        AC_MSG_ERROR([No Go compiler found.])
                    fi
                else
                    AC_MSG_RESULT([not found])
                    AC_MSG_ERROR([No Go compiler found.])
                fi
            fi
        fi
    fi
fi

if test -z "$GO_BIN_DIR"; then
    GO_BIN_DIR=`which "$GO_COMPILER" 2>/dev/null | xargs dirname 2>/dev/null || echo ""`
fi

if test "x$go_path" != "x"; then
    PATH="$save_PATH"
fi

# Detect Go compiler type
AC_MSG_CHECKING([for Go compiler type])
GO_VERSION_OUTPUT=`$GO_COMPILER version 2>/dev/null`
GO_COMPILER_TYPE=unknown

case "$GO_VERSION_OUTPUT" in
    *gccgo*)
        GO_COMPILER_TYPE=gccgo
        AC_DEFINE([HAVE_GCCGO], [1], [Define to 1 if using gccgo])
        ;;
    go*version*1.*)
        GO_COMPILER_TYPE=go
        AC_DEFINE([HAVE_GOOGLE_GO], [1], [Define to 1 if using Google Go])
        ;;
    *)
        GO_COMPILER_TYPE=unknown
        AC_MSG_WARN([Cannot determine Go compiler type from: $GO_VERSION_OUTPUT])
        ;;
esac
AC_MSG_RESULT([$GO_COMPILER_TYPE])
AC_SUBST([GO_COMPILER])

# Detect GOROOT
AC_MSG_CHECKING([for GOROOT])
if test -n "$GOROOT"; then
    GO_GOROOT="$GOROOT"
    AC_MSG_RESULT([$GO_GOROOT (from environment)])
elif test -n "$GO_BIN_DIR" -a -d "$GO_BIN_DIR"; then
    GO_GOROOT=`dirname "$GO_BIN_DIR" 2>/dev/null`
    if test -d "$GO_GOROOT/bin"; then
        AC_MSG_RESULT([$GO_GOROOT (from GO_BIN_DIR)])
    else
        GO_GOROOT=""
        AC_MSG_RESULT([not found (GO_BIN_DIR=$GO_BIN_DIR not inside a GOROOT)])
    fi
else
    GO_GOROOT=`"$GO_COMPILER" env GOROOT 2>/dev/null`
    if test -n "$GO_GOROOT" -a -d "$GO_GOROOT/bin"; then
        AC_MSG_RESULT([$GO_GOROOT (from 'go env GOROOT')])
    else
        GO_GOROOT=""
        AC_MSG_RESULT([not found])
        AC_MSG_WARN([GOROOT not set; CGo builds may fail with trimmed Go binaries])
    fi
fi
AC_SUBST([GO_GOROOT])

# Detect Go binary extension (".exe" on Windows)
AC_MSG_CHECKING([for Go binary extension])
GOOS=`$GO_COMPILER env GOOS 2>/dev/null`
if test "x$GOOS" = "xwindows"; then
    GO_EXEEXT=".exe"
else
    GO_EXEEXT=""
fi
AC_MSG_RESULT([$GO_EXEEXT])
AC_SUBST([GO_EXEEXT])

if test -n "$GO_BIN_DIR" && test -x "$GO_BIN_DIR/go$GO_EXEEXT"; then
    GO_COMPILER="$GO_BIN_DIR/go$GO_EXEEXT"
fi

AC_SUBST([GO_BIN_DIR])

# Generate output files
AC_CONFIG_FILES([Makefile])
AC_OUTPUT

# Summary
echo
echo "Build configuration summary:"
echo "  Package:        $PACKAGE_NAME $PACKAGE_VERSION"
echo "  Go compiler:    $GO_COMPILER"
echo "  GOROOT:         $GO_GOROOT"
echo "  GOOS:           $GOOS"
echo
