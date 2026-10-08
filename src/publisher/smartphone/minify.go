// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

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
// references a distinct pair of input/output files, so the same file can never
// be processed by more than one worker at a time.
type htMinifyJob struct {
	MinifyType string
	InFile     string
	OutFile    string
}

// registerMinifier registers the minifier matching the given mediatype on a
// minify.M instance. Only the relevant minifier is registered so that other
// minifiers never touch content of a different type.
func registerMinifier(m *minify.M, minifyType string) {
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

// runMinifyParallel runs the given jobs concurrently, one worker per available
// processor. Each worker owns its own minify.M instance (the minify library is
// not safe for concurrent use of a single instance), and the caller blocks
// until every job has finished. The first error (if any) is returned.
func runMinifyParallel(jobs []htMinifyJob) error {
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
					registerMinifier(m, job.MinifyType)
					ms[job.MinifyType] = m
				}
				if err := minifyCommonFile(m, job.MinifyType, job.InFile, job.OutFile); err != nil {
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

func minifyCommonFile(m *minify.M, minifyType string, inFile string, outFile string) error {
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

// HTMinifyFile minifies a single file from inFile into outFile using the
// mediatype minifier (for example "application/json", "text/css" or
// "application/javascript").
func HTMinifyFile(minifyType string, inFile string, outFile string) error {
	m := minify.New()
	registerMinifier(m, minifyType)
	return minifyCommonFile(m, minifyType, inFile, outFile)
}

func createDirectories(cfg Config, langs []string) error {
	if err := os.MkdirAll(cfg.ContentPath, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.ContentPath+"lang/", 0755); err != nil {
		return err
	}

	for _, lang := range langs {
		if err := os.MkdirAll(fmt.Sprintf("%slang/%s/", cfg.ContentPath, lang), 0755); err != nil {
			return err
		}
		if err := os.MkdirAll(fmt.Sprintf("%slang/%s/smartphone/", cfg.ContentPath, lang), 0755); err != nil {
			return err
		}
	}
	return nil
}

// removeOldContent deletes the previous output. It refuses to remove the
// current directory or a filesystem root to avoid accidental data loss when
// the content path is misconfigured.
func removeOldContent(cfg Config) error {
	clean := filepath.Clean(strings.TrimSpace(cfg.ContentPath))
	if clean == "." || clean == string(filepath.Separator) || clean == "" {
		return fmt.Errorf("refusing to remove content path: %s", cfg.ContentPath)
	}

	return os.RemoveAll(cfg.ContentPath)
}

// rewriteAndMinify transforms every smartphone file of a language into a
// temporary file and minifies it into the output directory.
func (r *runner) rewriteAndMinify(lang string) error {
	srcDir := fmt.Sprintf("%ssrc/smartphone/%s/", r.cfg.SrcPath, lang)
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		if r.cfg.Verbose {
			fmt.Println("Skipping missing smartphone directory", srcDir)
		}
		return nil
	}

	outDir := fmt.Sprintf("%slang/%s/smartphone/", r.cfg.ContentPath, lang)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	var jobs []htMinifyJob
	var tmpFiles []string
	failed := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		smGameFile := srcDir + entry.Name()
		tmpFile, err := r.transform(lang, smGameFile)
		if err != nil {
			failed++
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

	if err := runMinifyParallel(jobs); err != nil {
		return err
	}

	for _, tmpFile := range tmpFiles {
		if err := os.Remove(tmpFile); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR removing tmp %s: %v\n", tmpFile, err)
		}
	}

	if failed > 0 {
		return fmt.Errorf("%d smartphone file(s) failed to transform", failed)
	}
	return nil
}

// minifyAll regenerates the minified smartphone content of every language into
// the content directory.
func (r *runner) minifyAll() error {
	langs := discoverLangs(r.cfg)

	if err := removeOldContent(r.cfg); err != nil {
		return err
	}
	if err := createDirectories(r.cfg, langs); err != nil {
		return err
	}

	for _, lang := range langs {
		if err := r.rewriteAndMinify(lang); err != nil {
			return err
		}
	}

	if r.cfg.Verbose {
		fmt.Println("Completed successfully!")
	}
	return nil
}
