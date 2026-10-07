// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
	"strings"

	"historytracers-publisher/smartphone"
)

var (
	MinifyFlag          bool
	ValidateFlag        bool
	verboseFlag         bool
	ShowCompilationFlag bool
	logFileFlag         string

	cfg smartphone.Config
)

const (
	defaultSrcPath     = "."
	defaultContentPath = "build/www/"
)

// HTParseArg registers and parses the command line options.
func HTParseArg() {
	flag.BoolVar(&MinifyFlag, "minify", false, "Rewrite and minify all smartphone JSON files into the content directory. (default: false)")
	flag.BoolVar(&ValidateFlag, "validate", false, "Validate smartphone JSON files without producing output. (default: false)")
	flag.BoolVar(&verboseFlag, "verbose", false, "Print information messages during file processing. (default: false)")
	flag.BoolVar(&ShowCompilationFlag, "compilation", false, "Show the directories used by this build. (default: false)")

	flag.StringVar(&cfg.SrcPath, "src", defaultSrcPath, "Directory containing the source files (the repository root). (default: .)")
	flag.StringVar(&cfg.ContentPath, "www", defaultContentPath, "Directory for the generated content. (default: build/www/)")
	flag.StringVar(&logFileFlag, "logfile", "", "Path to a log file (truncates on open). All output is redirected here.")
	flag.StringVar(&cfg.DBPath, "db", "", "Path to the source SQLite database. (default: <src>/lang/sources/history_tracers.db when it exists)")

	flag.Parse()

	cfg.Verbose = verboseFlag
}

func htPrintOptions() {
	c := cfg.Normalized()
	fmt.Println("History Tracers smartphone publisher\n\nSource Path: ", strings.TrimSpace(c.SrcPath),
		"\nContent Path:", strings.TrimSpace(c.ContentPath),
		"\nSource DB:   ", strings.TrimSpace(c.SourceDBPath()))
}
