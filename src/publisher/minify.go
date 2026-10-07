// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	mjson "github.com/tdewolff/minify/v2/json"
)

// htMinifyJob is a single, self-contained file minification task. Each job
// references a distinct pair of input/output files, so the same file can
// never be processed by more than one worker at a time.
type htMinifyJob struct {
	MinifyType string
	InFile     string
	OutFile    string
}

// htRegisterMinifier registers the minifier matching the given mediatype on a
// minify.M instance. Only the relevant minifier is registered so that other
// minifiers never touch content of a different type.
func htRegisterMinifier(m *minify.M, minifyType string) {
	switch minifyType {
	case "application/json":
		m.AddFunc("application/json", mjson.Minify)
	case "application/javascript":
		m.AddFunc("application/javascript", js.Minify)
	case "text/css":
		m.AddFunc("text/css", css.Minify)
	case "text/html":
		m.AddFunc("text/html", html.Minify)
	}
}

// htRunMinifyParallel runs the given jobs concurrently, one worker per
// available processor. Each worker owns its own minify.M instance (the minify
// library is not safe for concurrent use of a single instance), and the
// caller blocks until every job has finished. The first error (if any) is
// returned.
func htRunMinifyParallel(jobs []htMinifyJob) error {
	if len(jobs) == 0 {
		return nil
	}

	numWorkers := runtime.NumCPU()
	if numWorkers > len(jobs) {
		numWorkers = len(jobs)
	}

	jobCh := make(chan htMinifyJob)
	errCh := make(chan error, len(jobs))

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ms := make(map[string]*minify.M)
			for job := range jobCh {
				m := ms[job.MinifyType]
				if m == nil {
					m = minify.New()
					htRegisterMinifier(m, job.MinifyType)
					ms[job.MinifyType] = m
				}
				if err := htMinifyCommonFile(m, job.MinifyType, job.InFile, job.OutFile); err != nil {
					errCh <- err
				}
			}
		}()
	}

	for _, job := range jobs {
		jobCh <- job
	}
	close(jobCh)
	wg.Wait()
	close(errCh)

	for err := range errCh {
		return err
	}
	return nil
}

func htMinifyCommonFile(m *minify.M, minifyType string, inFile string, outFile string) error {
	r, err1 := os.Open(inFile)
	if err1 != nil {
		return err1
	}
	defer r.Close()

	w, err2 := os.Create(outFile)
	if err2 != nil {
		return err2
	}
	defer w.Close()

	if err3 := m.Minify(minifyType, w, r); err3 != nil {
		return err3
	}

	return w.Close()
}

func htMinifyCreateDirectories() {
	htCreateDirectory(CFG.ContentPath)
	htCreateDirectory(CFG.ContentPath + "lang/")

	for _, lang := range htLangPaths {
		htCreateDirectory(fmt.Sprintf("%slang/%s/", CFG.ContentPath, lang))
		htCreateDirectory(fmt.Sprintf("%slang/%s/smartphone/", CFG.ContentPath, lang))
	}
}

// htMinifyRemoveOldContent deletes the previous output. It refuses to remove
// the current directory or a filesystem root to avoid accidental data loss
// when -www is misconfigured.
func htMinifyRemoveOldContent() {
	clean := filepath.Clean(strings.TrimSpace(CFG.ContentPath))
	if clean == "." || clean == string(filepath.Separator) || clean == "" {
		panic("refusing to remove content path: " + CFG.ContentPath)
	}

	if err := os.RemoveAll(CFG.ContentPath); err != nil {
		panic(err)
	}
}

// htRewriteAndMinifySmartphone transforms every smartphone file of a language
// into a temporary file and minifies it into the output directory.
func htRewriteAndMinifySmartphone(lang string) {
	srcDir := fmt.Sprintf("%ssrc/smartphone/%s/", CFG.SrcPath, lang)
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		if verboseFlag {
			fmt.Println("Skipping missing smartphone directory", srcDir)
		}
		return
	}

	outDir := fmt.Sprintf("%slang/%s/smartphone/", CFG.ContentPath, lang)

	var jobs []htMinifyJob
	var tmpFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		smGameFile := srcDir + entry.Name()
		tmpFile, err := htTransformSMGame(lang, smGameFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR transforming smartphone:", err)
			continue
		}

		tmpFiles = append(tmpFiles, tmpFile)
		jobs = append(jobs, htMinifyJob{
			MinifyType: "application/json",
			InFile:     tmpFile,
			OutFile:    outDir + entry.Name(),
		})
	}

	if err := htRunMinifyParallel(jobs); err != nil {
		panic(err)
	}

	for _, tmpFile := range tmpFiles {
		if err := os.Remove(tmpFile); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR removing tmp %s: %v\n", tmpFile, err)
		}
	}
}

// HTMinifyAllFiles regenerates the minified smartphone content of every
// language into the content directory.
func HTMinifyAllFiles() {
	htMinifyRemoveOldContent()
	htMinifyCreateDirectories()

	for _, lang := range htLangPaths {
		htRewriteAndMinifySmartphone(lang)
	}

	if verboseFlag {
		fmt.Println("Completed successfully!")
	}
}
