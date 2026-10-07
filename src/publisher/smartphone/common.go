// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/google/uuid"
)

const defaultLangPaths = "en-US,es-ES,pt-BR"

// discoverLangs returns the directories found in src/smartphone/. When the
// directory cannot be read, the well known History Tracers languages are used
// as a fallback.
func discoverLangs(cfg Config) []string {
	dir := cfg.SrcPath + "src/smartphone/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		return strings.Split(defaultLangPaths, ",")
	}

	var langs []string
	for _, entry := range entries {
		if entry.IsDir() {
			langs = append(langs, entry.Name())
		}
	}

	if len(langs) == 0 {
		return strings.Split(defaultLangPaths, ",")
	}

	sort.Strings(langs)
	return langs
}

// HTLanguages returns the language directories found under src/smartphone/.
func HTLanguages(cfg Config) []string {
	return discoverLangs(cfg.Normalized())
}

func htOpenFileReadClose(fileName string) ([]byte, error) {
	contentFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(contentFile)
	if err != nil {
		contentFile.Close()
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

	n, err := io.Copy(dfp, sfp)
	if n == 0 || err != nil {
		return err
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

// writeSmartphoneTmpFile writes data as indented JSON to a temporary file in
// the smartphone output directory. The returned path is used as the input of
// the minifier and removed once the file has been processed.
func writeSmartphoneTmpFile(cfg Config, lang string, data interface{}) (string, error) {
	id := uuid.New()
	tmpFile := fmt.Sprintf("%slang/%s/smartphone/%s.tmp", cfg.ContentPath, lang, id.String())

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
