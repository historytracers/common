// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	. "github.com/historytracers/common"
)

// htTransformSMGame reads a smartphone (or smGame) file, applies the content
// transformations used during publishing, and writes the result to a
// temporary file. The source file is never modified, so the generated output
// can be regenerated at any time without touching the repository content.
func htTransformSMGame(lang string, smGameFile string) (string, error) {
	if verboseFlag {
		fmt.Println("Adjusting file", smGameFile)
	}

	byteValue, err := htOpenFileReadClose(smGameFile)
	if err != nil {
		return "", err
	}

	var localSMGameFile SMGameFile
	err = json.Unmarshal(byteValue, &localSMGameFile)
	if err != nil {
		htCommonJSONError(byteValue, err)
		return "", err
	}

	if localSMGameFile.Sources != nil {
		htLoadSourceFromFile(localSMGameFile.Sources)
	}

	if len(localSMGameFile.Content) > 0 && localSMGameFile.Content[0].Smile == "" {
		localSMGameFile.Content[0].Smile = "nerd"
	}

	if len(localSMGameFile.Content) > 0 && localSMGameFile.Content[len(localSMGameFile.Content)-1].Smile == "" {
		localSMGameFile.Content[len(localSMGameFile.Content)-1].Smile = "party"
	}

	for i := range localSMGameFile.Content {
		if localSMGameFile.Content[i].Answer != nil && localSMGameFile.Content[i].Smile == "" {
			localSMGameFile.Content[i].Smile = "thinking"
		}
		for j := range localSMGameFile.Content[i].SourceMenu {
			if strings.HasPrefix(localSMGameFile.Content[i].SourceMenu[j].Page, "index.html?") {
				localSMGameFile.Content[i].SourceMenu[j].Page = "https://www.historytracers.org/" + localSMGameFile.Content[i].SourceMenu[j].Page
			}
		}
	}

	if _, fileWasModified := htGitModifiedMap[smGameFile]; fileWasModified && len(localSMGameFile.LastUpdate) > 0 {
		localSMGameFile.LastUpdate[0] = HTUpdateTimestamp()
	}

	return htWriteSmartphoneTmpFile(lang, &localSMGameFile)
}

// htValidateSMGameIDs checks that every content block has a valid UUID "id".
func htValidateSMGameIDs(smGameFile string) error {
	byteValue, err := htOpenFileReadClose(smGameFile)
	if err != nil {
		return err
	}

	var localSMGameFile SMGameFile
	err = json.Unmarshal(byteValue, &localSMGameFile)
	if err != nil {
		htCommonJSONError(byteValue, err)
		return fmt.Errorf("%s: unable to parse file as an SM file: %v", smGameFile, err)
	}

	if localSMGameFile.Type != "sm_game" {
		return nil
	}

	for i, block := range localSMGameFile.Content {
		if _, err := uuid.Parse(block.ID); err != nil {
			return fmt.Errorf("%s: content block %d has an invalid \"id\" field (expected uuid format, got %q)", smGameFile, i, block.ID)
		}
	}

	return nil
}

// htValidateSMGameFormats validates every smartphone JSON file found in the
// repository and returns the number of files that failed validation.
func htValidateSMGameFormats() int {
	invalid := 0
	for _, lang := range htLangPaths {
		smGameDir := fmt.Sprintf("%ssrc/smartphone/%s/", CFG.SrcPath, lang)
		entries, err := os.ReadDir(smGameDir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			if !strings.HasSuffix(entry.Name(), ".json") {
				continue
			}

			smGameFile := smGameDir + entry.Name()
			if err := htValidateSMGameIDs(smGameFile); err != nil {
				invalid++
				fmt.Fprintln(os.Stderr, "ERROR:", err)
			} else if verboseFlag {
				fmt.Println("OK:", smGameFile)
			}
		}
	}
	return invalid
}
