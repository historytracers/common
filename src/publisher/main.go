// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"os"
)

// htCreateDirectory creates a directory and any missing parents.
func htCreateDirectory(name string) {
	if err := os.MkdirAll(name, 0755); err != nil {
		panic(err)
	}
}

func htRunStopFlags() {
	htFillModifiedGit()

	var stopRun bool

	if ShowCompilationFlag {
		htPrintOptions()
		os.Exit(0)
	}

	if ValidateFlag {
		if invalid := htValidateSMGameFormats(); invalid > 0 {
			fmt.Fprintf(os.Stderr, "%d smartphone file(s) failed validation\n", invalid)
			os.Exit(1)
		}
		stopRun = true
	}

	if MinifyFlag {
		HTMinifyAllFiles()
		stopRun = true
	}

	if stopRun {
		os.Exit(0)
	}
}

func main() {
	HTParseArg()
	htInitializeCommonMaps()
	htLoadSmartphoneLangs()

	if logFileFlag != "" {
		f, err := os.Create(logFileFlag)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		os.Stdout = f
		os.Stderr = f
	}

	htRunStopFlags()

	fmt.Println("No action specified. Use --help to see available options.")
}
