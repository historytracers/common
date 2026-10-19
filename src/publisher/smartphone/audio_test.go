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

// writeTestMathKeywords writes a math_keywords.json file with the values that
// the audio code reads (times, plus, equals and divided by).
func writeTestMathKeywords(t *testing.T, root string, lang string, times string, plus string, equals string, dividedBy string) {
	t.Helper()

	keywords := make([]string, 37)
	keywords[33] = times
	keywords[34] = plus
	keywords[35] = equals
	keywords[36] = dividedBy

	payload, err := json.Marshal(common.HTMathKeywordsFormat{Keywords: keywords})
	if err != nil {
		t.Fatal(err)
	}

	dir := filepath.Join(root, "lang", lang)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "math_keywords.json"), payload, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestHTReplaceMultiplication(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "portuguese word xadrez", in: "o xadrez e a matemática", want: "o xadrez e a matemática"},
		{name: "portuguese word aux", in: "L’Anse aux Meadows", want: "L’Anse aux Meadows"},
		{name: "axis label", in: "eixo x", want: "eixo x"},
		{name: "name with x", in: "Yax Pasaj Chan Yopaat", want: "Yax Pasaj Chan Yopaat"},
		{name: "letter list", in: "w, x, y, z", want: "w, x, y, z"},
		{name: "numbers", in: "25 x 25", want: "25 vezes 25"},
		{name: "numbers without spaces", in: "5x6", want: "5 vezes 6"},
		{name: "variables", in: "a x N", want: "a vezes N"},
		{name: "fraction", in: "3×1/2", want: "3 vezes 1/2"},
		{name: "parentheses", in: "(a x b)", want: "(a vezes b)"},
		{name: "mixed operands", in: "2 x (3 + 4)", want: "2 vezes (3 + 4)"},
		{name: "function", in: "c × cos (θ)", want: "c vezes cos (θ)"},
		{name: "chained x", in: "2 x 3 x 4", want: "2 vezes 3 vezes 4"},
		{name: "chained sign", in: "2 × 3 × 4", want: "2 vezes 3 vezes 4"},
		{name: "word operands", in: "Total Fruits × Fingers", want: "Total Fruits vezes Fingers"},
		{name: "prose sign with comma", in: "o sinal de multiplicação ×, que lemos", want: "o sinal de multiplicação ×, que lemos"},
		{name: "heading sign", in: "Um novo sinal: × (vezes)", want: "Um novo sinal: × (vezes)"},
		{name: "word containing x before sign", in: "box x 3", want: "box x 3"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := htReplaceMultiplication(tc.in, " vezes "); got != tc.want {
				t.Errorf("htReplaceMultiplication(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestHTReplaceSubtraction(t *testing.T) {
	cases := []struct {
		name string
		lang string
		in   string
		want string
	}{
		{name: "ascii hyphen in equation", lang: "en-US", in: "9 - 1 = 8", want: "9 minus 1 = 8"},
		{name: "unicode minus in equation", lang: "en-US", in: "9 − 1 = 8", want: "9 minus 1 = 8"},
		{name: "unicode minus in prose expression", lang: "en-US", in: "the subtraction of 1000 − 175?", want: "the subtraction of 1000 minus 175?"},
		{name: "word operands in equation", lang: "en-US", in: "Protons in Tritium − Protons in Hydrogen = 1 − 1 = 0", want: "Protons in Tritium minus Protons in Hydrogen = 1 minus 1 = 0"},
		{name: "standalone equation before another equation line", lang: "en-US", in: "1000 - 175.\n\n900 + 100 - 175.", want: "1000 minus 175.\n\n900 + 100 minus 175."},
		{name: "hyphenated range", lang: "en-US", in: "581 - 618.", want: "581 - 618."},
		{name: "numeric range with words", lang: "en-US", in: "91 - 100 Metate simétrico.", want: "91 - 100 Metate simétrico."},
		{name: "minus sign in prose", lang: "en-US", in: "The minus sign (−) tells us the direction.", want: "The minus sign (−) tells us the direction."},
		{name: "button label", lang: "en-US", in: "Use plus and − to change the value.", want: "Use plus and − to change the value."},
		{name: "negative number", lang: "en-US", in: "-1 = -1", want: "-1 = -1"},
		{name: "portuguese", lang: "pt-BR", in: "9 - 1 = 8", want: "9 menos 1 = 8"},
		{name: "spanish", lang: "es-ES", in: "9 - 1 = 8", want: "9 menos 1 = 8"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := htReplaceSubtraction(tc.in, tc.lang); got != tc.want {
				t.Errorf("htReplaceSubtraction(%q, %q) = %q, want %q", tc.in, tc.lang, got, tc.want)
			}
		})
	}
}

func TestReplaceMathUsesLocalizedKeywords(t *testing.T) {
	cases := []struct {
		lang      string
		times     string
		plus      string
		equals    string
		wantTimes string
		wantMinus string
	}{
		{lang: "en-US", times: " times ", plus: " plus ", equals: " equals ", wantTimes: "2 times 3 plus 1 equals 7", wantMinus: "9 minus 1 equals 8"},
		{lang: "pt-BR", times: " vezes ", plus: " mais ", equals: " igual ", wantTimes: "2 vezes 3 mais 1 igual 7", wantMinus: "9 menos 1 igual 8"},
		{lang: "es-ES", times: " veces ", plus: " más ", equals: " igual ", wantTimes: "2 veces 3 más 1 igual 7", wantMinus: "9 menos 1 igual 8"},
	}

	for _, tc := range cases {
		t.Run(tc.lang, func(t *testing.T) {
			dir := t.TempDir()
			writeTestMathKeywords(t, dir, tc.lang, tc.times, tc.plus, tc.equals, "divided by")
			r := newRunner(Config{SrcPath: dir})

			if got := r.replaceMath("2 x 3 + 1 = 7", tc.lang); got != tc.wantTimes {
				t.Errorf("replaceMath multiplication = %q, want %q", got, tc.wantTimes)
			}
			if got := r.replaceMath("9 - 1 = 8", tc.lang); got != tc.wantMinus {
				t.Errorf("replaceMath subtraction = %q, want %q", got, tc.wantMinus)
			}
			if got := r.dividedByKeyword(tc.lang); got != "divided by" {
				t.Errorf("dividedByKeyword = %q, want %q", got, "divided by")
			}
		})
	}
}

func TestReplaceMathFallsBackToLanguageDefaults(t *testing.T) {
	r := newRunner(Config{SrcPath: t.TempDir()})

	if got := r.replaceMath("2 x 3 + 1 = 7", "en-US"); got != "2 times 3 plus 1 equals 7" {
		t.Errorf("replaceMath fallback = %q", got)
	}
	if got := r.dividedByKeyword("en-US"); got != "divided by" {
		t.Errorf("dividedByKeyword fallback = %q", got)
	}
	if got := r.replaceMath("2 x 3 + 1 = 7", "pt-BR"); got != "2 vezes 3 mais 1 igual 7" {
		t.Errorf("replaceMath pt-BR fallback = %q", got)
	}
	if got := r.dividedByKeyword("pt-BR"); got != "dividido por" {
		t.Errorf("dividedByKeyword pt-BR fallback = %q", got)
	}
}

func TestHTRemoveHTMLTags(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "tags", in: "<p>Hello <b>world</b></p>", want: "Hello world"},
		{name: "comparison symbols", in: "0 < 1 and 2 > 1", want: "0 < 1 and 2 > 1"},
		{name: "symbol in parentheses", in: "We use the symbol (<), so 20 < 21.", want: "We use the symbol (<), so 20 < 21."},
		{name: "svg", in: "a<svg><text>x</text></svg>b", want: "ab"},
		{name: "style", in: "a<style>.x{color:red}</style>b", want: "ab"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := htRemoveHTMLTags(tc.in); got != tc.want {
				t.Errorf("htRemoveHTMLTags(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestHTIsInternalContentLink(t *testing.T) {
	cases := []struct {
		href string
		want bool
	}{
		{href: "index.html?page=class_content&arg=first_steps_menu", want: true},
		{href: "./index.html?page=physics", want: true},
		{href: "../index.html", want: true},
		{href: "/index.html?page=first_steps", want: true},
		{href: "https://historytracers.org/index.html?page=physics", want: true},
		{href: "https://www.historytracers.org/index.html?page=families", want: true},
		{href: "https://example.com/index.html", want: false},
		{href: "https://historytracers.org/images/logo.png", want: false},
		{href: "mailto:contact@historytracers.org", want: false},
		{href: "", want: false},
	}

	for _, tc := range cases {
		if got := htIsInternalContentLink(tc.href); got != tc.want {
			t.Errorf("htIsInternalContentLink(%q) = %v, want %v", tc.href, got, tc.want)
		}
	}
}

func TestHTReplaceMarkdownLinks(t *testing.T) {
	internal := htReplaceMarkdownLinks("Read [First Steps](index.html?page=first_steps_menu) now.")
	if internal != "Read First Steps now." {
		t.Errorf("internal link = %q", internal)
	}

	external := "[Example](https://example.com/page)"
	if got := htReplaceMarkdownLinks(external); got != external {
		t.Errorf("external link was changed: %q", got)
	}

	image := "![Figure](index.html)"
	if got := htReplaceMarkdownLinks(image); got != image {
		t.Errorf("image was changed: %q", got)
	}
}

func TestHTReplaceFractions(t *testing.T) {
	cases := []struct {
		name      string
		in        string
		dividedBy string
		want      string
	}{
		{name: "parenthesized fraction", in: "one over two (1/2)", dividedBy: "divided by", want: "one over two (1 divided by 2)"},
		{name: "repeated fraction", in: "Each finger is 1/10.", dividedBy: "divided by", want: "Each finger is 1 divided by 10."},
		{name: "whole fraction", in: "gives 10/10, that is 1", dividedBy: "divided by", want: "gives 10 divided by 10, that is 1"},
		{name: "portuguese", in: "um sobre dois (1/2)", dividedBy: "dividido por", want: "um sobre dois (1 dividido por 2)"},
		{name: "fraction slash", in: "half (1⁄2)", dividedBy: "divided by", want: "half (1 divided by 2)"},
		{name: "fraction entity", in: "half (1&frasl;2)", dividedBy: "divided by", want: "half (1 divided by 2)"},
		{name: "url is untouched", in: "see https://example.com/2024/10 now", dividedBy: "divided by", want: "see https://example.com/2024/10 now"},
		{name: "ratio is untouched", in: "the scale is 1:100", dividedBy: "divided by", want: "the scale is 1:100"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := htReplaceFractions(tc.in, tc.dividedBy); got != tc.want {
				t.Errorf("htReplaceFractions(%q, %q) = %q, want %q", tc.in, tc.dividedBy, got, tc.want)
			}
		})
	}
}

func TestAdjustAudioStringBeforeWriteKeepsMathSigns(t *testing.T) {
	dir := t.TempDir()
	writeTestMathKeywords(t, dir, "pt-BR", " vezes ", " mais ", " igual ", "dividido por")
	r := newRunner(Config{SrcPath: dir})

	in := "O xadrez e 2 x 3. O sinal de multiplicação ×, que lemos. 3 − 1 = 2. A metade (1/2)."
	got := r.adjustAudioStringBeforeWrite(in, "pt-BR")

	for _, want := range []string{"O xadrez", "2 vezes 3", "multiplicação ×", "3 menos 1 igual 2", "1 dividido por 2"} {
		if !strings.Contains(got, want) {
			t.Errorf("adjustAudioStringBeforeWrite(%q) = %q, missing %q", in, got, want)
		}
	}
	if strings.Contains(got, "vezes adrez") {
		t.Errorf("word was corrupted: %q", got)
	}
}

func TestHTGenerateAudioKeepsComparisonSymbols(t *testing.T) {
	dir := t.TempDir()
	game := &common.SMGameFile{
		Title:      "Test",
		Type:       "sm_game",
		LastUpdate: []string{"0"},
		Content: []common.SMGameContent{
			{
				ID: screenOneID,
				Text: []common.HTText{
					{
						Text:   "#### Comparing\n\nWe use the symbol (<), so 20 < 21. Also 0 < 1 and 2 > 1.",
						Format: "markdown",
					},
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
	for _, want := range []string{"symbol (<)", "20 < 21", "0 < 1", "2 > 1"} {
		if !strings.Contains(content, want) {
			t.Errorf("audio content %q is missing %q", content, want)
		}
	}
}
