// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "modernc.org/sqlite"

	"github.com/historytracers/common"
)

// sourceMap holds the sources referenced by the file currently being
// processed. allSourceMap keeps every source seen so far, which makes it
// possible to detect duplicated UUIDs with diverging fields.
var sourceMap map[string]common.HTSourceElement
var allSourceMap map[string]common.HTSourceElement

// sourceDBMissingWarned ensures the "database not found" message is printed at
// most once per run.
var sourceDBMissingWarned bool

func htInitializeCommonMaps() {
	sourceMap = make(map[string]common.HTSourceElement)
	allSourceMap = make(map[string]common.HTSourceElement)
}

func htCompareSources(first *common.HTSourceElement, second *common.HTSourceElement) bool {
	if first.ID == second.ID &&
		first.Citation == second.Citation &&
		first.PublishDate == second.PublishDate &&
		first.URL == second.URL {
		return true
	}

	return false
}

func htFillSourceMap(src []common.HTSourceElement, fileID string) {
	for _, element := range src {
		if _, ok := sourceMap[element.ID]; !ok {
			sourceMap[element.ID] = element
		}

		if stored, ok := allSourceMap[element.ID]; !ok {
			allSourceMap[element.ID] = element
		} else if ok {
			if !htCompareSources(&stored, &element) {
				fmt.Fprintf(os.Stderr, "Duplicate UUID %s: MISSING fields in one entry (source file: %s).\n", element.ID, fileID)
			}
		}
	}
}

func htFillSourcesMap(src *common.HTSourceFile, fileID string) {
	if src.PrimarySources != nil {
		htFillSourceMap(src.PrimarySources, fileID)
	}
	if src.ReferencesSources != nil {
		htFillSourceMap(src.ReferencesSources, fileID)
	}
	if src.ReligiousSources != nil {
		htFillSourceMap(src.ReligiousSources, fileID)
	}
	if src.SocialMediaSources != nil {
		htFillSourceMap(src.SocialMediaSources, fileID)
	}
}

func htLoadSourceFromFile(srcs []string) {
	htLoadSourceFromDB(srcs)
}

// htLoadHTSourceFileFromDB reads every citation associated with the given file
// ID and groups it by citation type.
func htLoadHTSourceFileFromDB(db *sql.DB, fileID string) *common.HTSourceFile {
	rows, err := db.Query(`
		SELECT c.cit_type, s.src_id, COALESCE(s.sfo_id, ''), s.src_citation, s.src_date, s.src_publish_date, COALESCE(s.src_url, '')
		FROM citation c
		JOIN sources s ON c.src_id = s.src_id
		WHERE c.fil_id = ?
	`, fileID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR querying sources for %s: %v\n", fileID, err)
		return nil
	}
	defer rows.Close()

	sf := &common.HTSourceFile{
		License:    []string{"SPDX-License-Identifier: GPL-3.0-or-later", "CC BY-NC 4.0 DEED"},
		LastUpdate: []string{""},
		Version:    1,
		Type:       "sources",
	}

	for rows.Next() {
		var citType int
		var elem common.HTSourceElement
		if err := rows.Scan(&citType, &elem.ID, &elem.SfoID, &elem.Citation, &elem.Date, &elem.PublishDate, &elem.URL); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR scanning row for %s: %v\n", fileID, err)
			continue
		}
		switch citType {
		case 0:
			sf.PrimarySources = append(sf.PrimarySources, elem)
		case 1:
			sf.ReferencesSources = append(sf.ReferencesSources, elem)
		case 2:
			sf.ReligiousSources = append(sf.ReligiousSources, elem)
		case 3:
			sf.SocialMediaSources = append(sf.SocialMediaSources, elem)
		}
	}
	return sf
}

// htLoadSourceFromDB loads the given file IDs from the optional sources
// database. When the database is not available it silently returns, because
// the smartphone files already carry the citation data needed for publishing.
func htLoadSourceFromDB(srcs []string) {
	dbPath := htSourceDBPath()

	if !htSourceDBExists(dbPath) {
		if verboseFlag && !sourceDBMissingWarned {
			sourceDBMissingWarned = true
			fmt.Println("Source database not found, skipping source loading:", dbPath)
		}
		return
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR opening source database %s: %v\n", dbPath, err)
		return
	}
	defer db.Close()

	for _, ptr := range srcs {
		if strings.TrimSpace(ptr) == "" {
			continue
		}
		sf := htLoadHTSourceFileFromDB(db, ptr)
		if sf == nil {
			continue
		}
		htFillSourcesMap(sf, ptr)
	}
}
