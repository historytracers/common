// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "modernc.org/sqlite"

	"github.com/historytracers/common"
)

func (r *runner) compareSources(first *common.HTSourceElement, second *common.HTSourceElement) bool {
	if first.ID == second.ID &&
		first.Citation == second.Citation &&
		first.PublishDate == second.PublishDate &&
		first.URL == second.URL {
		return true
	}

	return false
}

func (r *runner) fillSourceMap(src []common.HTSourceElement, fileID string) {
	for _, element := range src {
		if _, ok := r.sourceMap[element.ID]; !ok {
			r.sourceMap[element.ID] = element
		}

		if stored, ok := r.allSourceMap[element.ID]; !ok {
			r.allSourceMap[element.ID] = element
		} else if ok {
			if !r.compareSources(&stored, &element) {
				fmt.Fprintf(os.Stderr, "Duplicate UUID %s: MISSING fields in one entry (source file: %s).\n", element.ID, fileID)
			}
		}
	}
}

func (r *runner) fillSourcesMap(src *common.HTSourceFile, fileID string) {
	if src.PrimarySources != nil {
		r.fillSourceMap(src.PrimarySources, fileID)
	}
	if src.ReferencesSources != nil {
		r.fillSourceMap(src.ReferencesSources, fileID)
	}
	if src.ReligiousSources != nil {
		r.fillSourceMap(src.ReligiousSources, fileID)
	}
	if src.SocialMediaSources != nil {
		r.fillSourceMap(src.SocialMediaSources, fileID)
	}
}

func (r *runner) loadSourceFromFile(srcs []string) {
	r.loadSourceFromDB(srcs)
}

// loadHTSourceFileFromDB reads every citation associated with the given file ID
// and groups it by citation type.
func loadHTSourceFileFromDB(db *sql.DB, fileID string) *common.HTSourceFile {
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

// loadSourceFromDB loads the given file IDs from the optional sources
// database. When the database is not available it silently returns, because
// the smartphone files already carry the citation data needed for publishing.
func (r *runner) loadSourceFromDB(srcs []string) {
	dbPath := r.cfg.SourceDBPath()

	if !sourceDBExists(dbPath) {
		if r.cfg.Verbose && !r.warnedDB {
			r.warnedDB = true
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
		sf := loadHTSourceFileFromDB(db, ptr)
		if sf == nil {
			continue
		}
		r.fillSourcesMap(sf, ptr)
	}
}
