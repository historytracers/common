// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/historytracers/common"
)

const (
	testLang     = "en-US"
	testGameID   = "11111111-1111-4111-8111-111111111111"
	screenOneID  = "22222222-2222-4222-8222-222222222222"
	screenTwoID  = "33333333-3333-4333-8333-333333333333"
	screenLastID = "44444444-4444-4444-8444-444444444444"
)

func writeTestGame(t *testing.T, dir string, game *common.SMGameFile) {
	t.Helper()

	srcDir := filepath.Join(dir, "src", "smartphone", testLang)
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatal(err)
	}

	bv, err := json.Marshal(game)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, testGameID+".json"), bv, 0644); err != nil {
		t.Fatal(err)
	}
}

// TestHTMinifyAllFiles exercises the reusable entry point: it writes a source
// file, rewrites and minifies it, and checks the applied transformations.
func TestHTMinifyAllFiles(t *testing.T) {
	dir := t.TempDir()
	game := &common.SMGameFile{
		Title:      "Test",
		Type:       "sm_game",
		LastUpdate: []string{"0"},
		Content: []common.SMGameContent{
			{
				ID:         screenOneID,
				SourceMenu: []common.HTSource{{Page: "index.html?page=class_content&arg=x"}},
			},
			{ID: screenTwoID, Answer: "yes"},
			{ID: screenLastID},
		},
	}
	writeTestGame(t, dir, game)

	cfg := Config{SrcPath: dir, ContentPath: filepath.Join(dir, "build", "www")}
	if err := HTMinifyAllFiles(cfg); err != nil {
		t.Fatalf("HTMinifyAllFiles: %v", err)
	}

	out := filepath.Join(dir, "build", "www", "lang", testLang, "smartphone", testGameID+".json")
	bv, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("reading output: %v", err)
	}

	var got common.SMGameFile
	if err := json.Unmarshal(bv, &got); err != nil {
		t.Fatalf("parsing output: %v", err)
	}

	if len(got.Content) != 3 {
		t.Fatalf("expected 3 screens, got %d", len(got.Content))
	}
	if got.Content[0].Smile != "nerd" {
		t.Errorf("first smile = %q, want nerd", got.Content[0].Smile)
	}
	if got.Content[1].Smile != "thinking" {
		t.Errorf("question smile = %q, want thinking", got.Content[1].Smile)
	}
	if got.Content[2].Smile != "party" {
		t.Errorf("last smile = %q, want party", got.Content[2].Smile)
	}
	wantPage := "https://www.historytracers.org/index.html?page=class_content&arg=x"
	if got.Content[0].SourceMenu[0].Page != wantPage {
		t.Errorf("source page = %q, want %q", got.Content[0].SourceMenu[0].Page, wantPage)
	}

	if invalid := HTValidateSMGameFormats(cfg); invalid != 0 {
		t.Errorf("HTValidateSMGameFormats = %d, want 0", invalid)
	}
}

// TestHTGenerateAudio checks that text-to-speech input files are generated in
// the configured audio directory.
func TestHTGenerateAudio(t *testing.T) {
	dir := t.TempDir()
	game := &common.SMGameFile{
		Title:      "Test",
		Type:       "sm_game",
		LastUpdate: []string{"0"},
		Content: []common.SMGameContent{
			{
				ID: screenOneID,
				Text: []common.HTText{
					{Text: "Hello **world**.", Format: "markdown"},
					{Text: "<p>Second paragraph.</p>", Format: "html"},
				},
			},
		},
	}
	writeTestGame(t, dir, game)

	cfg := Config{
		SrcPath:     dir,
		ContentPath: filepath.Join(dir, "build", "www"),
		AudioPath:   filepath.Join(dir, "audio"),
	}
	if err := HTGenerateAudio(cfg); err != nil {
		t.Fatalf("HTGenerateAudio: %v", err)
	}

	out := filepath.Join(dir, "audio", testGameID+"_"+screenOneID+"_"+testLang+".txt")
	bv, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("reading audio file: %v", err)
	}

	content := string(bv)
	if strings.Contains(content, "<p>") {
		t.Errorf("audio content still contains HTML tags: %q", content)
	}
	if !strings.Contains(content, "Hello world") {
		t.Errorf("audio content missing expected text: %q", content)
	}
	if !strings.Contains(content, "Second paragraph") {
		t.Errorf("audio content missing expected text: %q", content)
	}
}

// TestHTValidateSMGameFormats checks that invalid content is counted.
func TestHTValidateSMGameFormats(t *testing.T) {
	dir := t.TempDir()

	srcDir := filepath.Join(dir, "src", "smartphone", testLang)
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatal(err)
	}
	bad := []byte(`{"type":"sm_game","content":[{"id":"not-a-uuid"}]}`)
	if err := os.WriteFile(filepath.Join(srcDir, "bad.json"), bad, 0644); err != nil {
		t.Fatal(err)
	}

	cfg := Config{SrcPath: dir, ContentPath: filepath.Join(dir, "build", "www")}
	if invalid := HTValidateSMGameFormats(cfg); invalid != 1 {
		t.Fatalf("HTValidateSMGameFormats = %d, want 1", invalid)
	}
}
