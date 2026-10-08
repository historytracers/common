// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"os"

	"historytracers-publisher/smartphone"
)

func htRunStopFlags() {
	var stopRun bool

	if ShowCompilationFlag {
		htPrintOptions()
		os.Exit(0)
	}

	if ValidateFlag {
		if invalid := smartphone.HTValidateSMGameFormats(cfg); invalid > 0 {
			fmt.Fprintf(os.Stderr, "%d smartphone file(s) failed validation\n", invalid)
			os.Exit(1)
		}
		stopRun = true
	}

	if MinifyFlag {
		if err := smartphone.HTMinifyAllFiles(cfg); err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
		stopRun = true
	}

	if AudioFlag {
		if err := smartphone.HTGenerateAudio(cfg); err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}
		stopRun = true
	}

	if stopRun {
		os.Exit(0)
	}
}

func main() {
	HTParseArg()

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
