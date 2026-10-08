// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"github.com/historytracers/common"
)

// runner holds the configuration and the state shared while processing the
// smartphone files, so the exported helpers can stay free of global state.
type runner struct {
	cfg          Config
	gitModified  map[string]bool
	sourceMap    map[string]common.HTSourceElement
	allSourceMap map[string]common.HTSourceElement
	warnedDB     bool
}

func newRunner(cfg Config) *runner {
	cfg = cfg.Normalized()
	return &runner{
		cfg:          cfg,
		gitModified:  gitModified(cfg),
		sourceMap:    make(map[string]common.HTSourceElement),
		allSourceMap: make(map[string]common.HTSourceElement),
	}
}

// HTMinifyAllFiles rewrites (normalizes) and minifies every JSON file found in
// src/smartphone/<lang>/ into <ContentPath>/lang/<lang>/smartphone/. The
// source files are never modified.
//
// The previous content directory is removed first, so ContentPath must point
// to a generated location (build/www/ by default) and never to the repository
// root.
func HTMinifyAllFiles(cfg Config) error {
	return newRunner(cfg).minifyAll()
}

// HTRewriteAndMinifySmartphone rewrites and minifies the smartphone files of a
// single language. It does not remove previously generated content.
func HTRewriteAndMinifySmartphone(cfg Config, lang string) error {
	r := newRunner(cfg)
	if err := createDirectories(r.cfg, []string{lang}); err != nil {
		return err
	}
	return r.rewriteAndMinify(lang)
}

// HTTransformSMGame rewrites a single smartphone (or smGame) file and writes
// the result to a temporary file in the output directory. It returns the path
// of the temporary file; the caller is responsible for removing it.
func HTTransformSMGame(cfg Config, lang string, smGameFile string) (string, error) {
	r := newRunner(cfg)
	return r.transform(lang, smGameFile)
}

// HTValidateSMGameFormats validates every smartphone JSON file and returns the
// number of files that failed validation.
func HTValidateSMGameFormats(cfg Config) int {
	return newRunner(cfg).validate()
}
