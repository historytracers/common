// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/historytracers/common"
)

// htConfig carries the source and content directories used by the publisher.
// It embeds the shared configuration base so the paths can be reused by other
// History Tracers tools.
type htConfig struct {
	common.HTConfigBase
}

var (
	MinifyFlag          bool
	ValidateFlag        bool
	verboseFlag         bool
	ShowCompilationFlag bool
	srcVal              string
	contentVal          string
	logFileFlag         string
	dbFileFlag          string
	CFG                 *htConfig
)

const (
	defaultSrcPath     = "."
	defaultContentPath = "build/www/"
	defaultSourceDB    = "lang/sources/history_tracers.db"
)

// htNormalizeDir makes sure a directory path ends with a forward slash so it
// can be concatenated directly with relative file names. Forward slashes are
// used on every platform because Git and the Go standard library accept them
// on Windows as well.
func htNormalizeDir(path string) string {
	if len(path) == 0 {
		return path
	}
	if !strings.HasSuffix(path, "/") && !strings.HasSuffix(path, "\\") {
		path += "/"
	}
	return path
}

// HTParseArg registers and parses the command line options.
func HTParseArg() {
	flag.BoolVar(&MinifyFlag, "minify", false, "Rewrite and minify all smartphone JSON files into the content directory. (default: false)")
	flag.BoolVar(&ValidateFlag, "validate", false, "Validate smartphone JSON files without producing output. (default: false)")
	flag.BoolVar(&verboseFlag, "verbose", false, "Print information messages during file processing. (default: false)")
	flag.BoolVar(&ShowCompilationFlag, "compilation", false, "Show the directories used by this build. (default: false)")

	flag.StringVar(&srcVal, "src", defaultSrcPath, "Directory containing the source files (the repository root). (default: .)")
	flag.StringVar(&contentVal, "www", defaultContentPath, "Directory for the generated content. (default: build/www/)")
	flag.StringVar(&logFileFlag, "logfile", "", "Path to a log file (truncates on open). All output is redirected here.")
	flag.StringVar(&dbFileFlag, "db", "", "Path to the source SQLite database. (default: <src>/lang/sources/history_tracers.db when it exists)")

	flag.Parse()

	CFG = htCreateConfig()
}

func htCreateConfig() *htConfig {
	return &htConfig{
		HTConfigBase: *common.NewHTConfigBase(0, htNormalizeDir(srcVal), htNormalizeDir(contentVal), ""),
	}
}

// htSourceDBPath returns the location of the academic sources database. The
// database is optional: the publisher works without it, but when present it is
// used to load citation data referenced by the smartphone files.
func htSourceDBPath() string {
	if len(dbFileFlag) > 0 {
		return dbFileFlag
	}
	return CFG.SrcPath + defaultSourceDB
}

func htSourceDBExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func htPrintOptions() {
	fmt.Println("History Tracers smartphone publisher\n\nSource Path: ", strings.TrimSpace(CFG.SrcPath),
		"\nContent Path:", strings.TrimSpace(CFG.ContentPath),
		"\nSource DB:   ", strings.TrimSpace(htSourceDBPath()))
}
