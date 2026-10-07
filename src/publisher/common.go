// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/google/uuid"
)

// htLangPaths holds the language directories discovered under
// src/smartphone/. It is populated by htLoadSmartphoneLangs.
var htLangPaths []string

const defaultLangPaths = "en-US,es-ES,pt-BR"

// htLoadSmartphoneLangs fills htLangPaths with the directories found in
// src/smartphone/. When the directory cannot be read, the well known History
// Tracers languages are used as a fallback.
func htLoadSmartphoneLangs() {
	dir := CFG.SrcPath + "src/smartphone/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		htLangPaths = strings.Split(defaultLangPaths, ",")
		return
	}

	var langs []string
	for _, entry := range entries {
		if entry.IsDir() {
			langs = append(langs, entry.Name())
		}
	}

	if len(langs) == 0 {
		htLangPaths = strings.Split(defaultLangPaths, ",")
		return
	}

	sort.Strings(langs)
	htLangPaths = langs
}

func htOpenFileReadClose(fileName string) ([]byte, error) {
	contentFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(contentFile)
	if err != nil {
		return nil, err
	}
	contentFile.Close()

	return byteValue, nil
}

// HTCopyFilesWithoutChanges copies the contents of srcFile over dstFile.
func HTCopyFilesWithoutChanges(dstFile string, srcFile string) error {
	srcStat, err := os.Stat(srcFile)
	if err != nil {
		return err
	}

	if !srcStat.Mode().IsRegular() {
		return nil
	}

	sfp, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer sfp.Close()

	dfp, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer dfp.Close()

	bytes, err := io.Copy(dfp, sfp)
	if bytes == 0 || err != nil {
		return err
	}

	if verboseFlag {
		fmt.Println("Copying file", srcFile, " to ", dstFile)
	}
	return nil
}

// htCommonJSONError reports the location of a JSON parsing error.
func htCommonJSONError(byteValue []byte, err error) {
	switch t := err.(type) {
	case *json.SyntaxError:
		begin := t.Offset
		if begin > 256 {
			begin -= 256
		} else if begin > 30 {
			begin -= 30
		}
		jsn := string(byteValue[begin:t.Offset])
		jsn += "<--(Invalid Character)"
		fmt.Fprintf(os.Stderr, "Invalid character at offset %v\n %s", t.Offset, jsn)
	case *json.UnmarshalTypeError:
		begin := t.Offset
		if begin > 256 {
			begin -= 256
		} else if begin > 30 {
			begin -= 30
		}
		jsn := string(byteValue[begin:t.Offset])
		jsn += "<--(Invalid Type)"
		fmt.Fprintf(os.Stderr, "Invalid type at offset %v\n %s", t.Offset, jsn)
	default:
		fmt.Printf("Invalid character at offset\n %s", err.Error())
	}
}

// htWriteSmartphoneTmpFile writes data as indented JSON to a temporary file in
// the smartphone output directory. The returned path is used as the input of
// the minifier and removed once the file has been processed.
func htWriteSmartphoneTmpFile(lang string, data interface{}) (string, error) {
	id := uuid.New()
	tmpFile := fmt.Sprintf("%slang/%s/smartphone/%s.tmp", CFG.ContentPath, lang, id.String())

	fp, err := os.Create(tmpFile)
	if err != nil {
		return "", err
	}

	e := json.NewEncoder(fp)
	e.SetEscapeHTML(false)
	e.SetIndent("", "   ")
	e.Encode(data)

	fp.Close()

	return tmpFile, nil
}
