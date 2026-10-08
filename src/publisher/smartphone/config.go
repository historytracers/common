// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"os"
	"strings"
)

// Config controls how the smartphone content is rewritten and minified.
type Config struct {
	// SrcPath is the repository root containing src/smartphone/. When empty,
	// the current directory is used.
	SrcPath string
	// ContentPath is the directory that receives the generated JSON. When
	// empty, build/www/ is used.
	ContentPath string
	// DBPath is the optional SQLite database with the academic sources. When
	// empty, <SrcPath>/lang/sources/history_tracers.db is used.
	DBPath string
	// AudioPath is the directory that receives the generated text-to-speech
	// input files. When empty, audio/ is used.
	AudioPath string
	// Verbose enables informational messages.
	Verbose bool
}

const (
	defaultSrcPath     = "./"
	defaultContentPath = "build/www/"
	defaultAudioPath   = "audio/"
	defaultSourceDB    = "lang/sources/history_tracers.db"
)

// normalizeDir makes sure a directory path ends with a forward slash so it can
// be concatenated directly with relative file names. Forward slashes are used
// on every platform because Git and the Go standard library accept them on
// Windows as well.
func normalizeDir(path string) string {
	if len(path) == 0 {
		return path
	}
	if !strings.HasSuffix(path, "/") && !strings.HasSuffix(path, "\\") {
		path += "/"
	}
	return path
}

// Normalized returns a copy of the configuration with the default directories
// applied and the paths guaranteed to end with a separator.
func (cfg Config) Normalized() Config {
	cfg.SrcPath = normalizeDir(cfg.SrcPath)
	if cfg.SrcPath == "" {
		cfg.SrcPath = defaultSrcPath
	}
	cfg.ContentPath = normalizeDir(cfg.ContentPath)
	if cfg.ContentPath == "" {
		cfg.ContentPath = defaultContentPath
	}
	cfg.AudioPath = normalizeDir(cfg.AudioPath)
	if cfg.AudioPath == "" {
		cfg.AudioPath = defaultAudioPath
	}
	return cfg
}

// SourceDBPath returns the location of the academic sources database. The
// database is optional: the publisher works without it, but when present it is
// used to load the citation data referenced by the smartphone files.
func (cfg Config) SourceDBPath() string {
	c := cfg.Normalized()
	if len(c.DBPath) > 0 {
		return c.DBPath
	}
	return c.SrcPath + defaultSourceDB
}

func sourceDBExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
