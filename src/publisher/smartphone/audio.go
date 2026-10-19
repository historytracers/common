// SPDX-License-Identifier: GPL-3.0-or-later

package smartphone

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/historytracers/common"
)

// HTGenerateAudio converts the text of every smartphone screen into plain text
// files suitable for a text-to-speech engine (for example Piper). The files are
// written to <AudioPath>/<lesson-uuid>_<screen-uuid>_<lang>.txt.
func HTGenerateAudio(cfg Config) error {
	return newRunner(cfg).generateAudio()
}

// HTGenerateAudioForLang generates the audio text files of a single language.
func HTGenerateAudioForLang(cfg Config, lang string) error {
	return newRunner(cfg).generateAudioLang(lang)
}

func (r *runner) generateAudio() error {
	for _, lang := range discoverLangs(r.cfg) {
		if err := r.generateAudioLang(lang); err != nil {
			return err
		}
	}
	return nil
}

func (r *runner) generateAudioLang(lang string) error {
	smGameDir := fmt.Sprintf("%ssrc/smartphone/%s/", r.cfg.SrcPath, lang)
	files, err := os.ReadDir(smGameDir)
	if err != nil {
		if r.cfg.Verbose {
			fmt.Println("Skipping missing smartphone directory", smGameDir)
		}
		return nil
	}

	if err := os.MkdirAll(r.cfg.AudioPath, 0755); err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".json") {
			continue
		}

		filePath := smGameDir + fileName
		byteValue, err := htOpenFileReadClose(filePath)
		if err != nil {
			continue
		}

		var smGame common.SMGameFile
		if err := json.Unmarshal(byteValue, &smGame); err != nil || smGame.Type != "sm_game" {
			continue
		}

		baseName := strings.TrimSuffix(fileName, ".json")
		for _, block := range smGame.Content {
			audioContent := htBuildSMGameAudioText(block.Text)
			if len(audioContent) == 0 {
				continue
			}

			audioContent = r.adjustAudioStringBeforeWrite(audioContent, lang)
			audioContent = htRemoveAsiaticCharacters(audioContent)

			if err := r.writeAudioFile(baseName+"_"+block.ID, lang, audioContent); err != nil {
				if r.cfg.Verbose {
					fmt.Fprintln(os.Stderr, "ERROR writing audio for", baseName+"_"+block.ID+":", err)
				}
				continue
			}
		}
	}
	return nil
}

func (r *runner) writeAudioFile(fileName string, lang string, content string) error {
	localPath := fmt.Sprintf("%s%s_%s.txt", r.cfg.AudioPath, fileName, lang)

	fp, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer fp.Close()

	content = htAdjustTrailingDots(content)
	if _, err := fp.WriteString(content); err != nil {
		return err
	}

	if r.cfg.Verbose {
		fmt.Println("Writing audio file", localPath)
	}
	return nil
}

func htBuildSMGameAudioText(textBlocks []common.HTText) string {
	var audioBuilder strings.Builder

	for _, textBlock := range textBlocks {
		if textBlock.Text == "" {
			continue
		}

		cleanText := htRemoveHTMLTags(textBlock.Text)
		cleanText = htRemoveAsiaticCharacters(cleanText)
		cleanText = strings.TrimSpace(cleanText)

		if len(cleanText) > 0 {
			audioBuilder.WriteString(cleanText)
			audioBuilder.WriteString("\n")
		}
	}

	return audioBuilder.String()
}

func (r *runner) adjustAudioStringBeforeWrite(str string, lang string) string {
	tableLineRegex := regexp.MustCompile(`(?m)^\+\-+(?:\+\-+)+\+$`)
	dashLineRegex := regexp.MustCompile(`^\s*-+\s*$`)
	patternLinksRegex := regexp.MustCompile(`\(\s*(?:;+\s*)+\)`)

	lines := strings.Split(str, "\n")
	final := ""
	for _, line := range lines {
		if tableLineRegex.MatchString(line) {
			continue
		}

		if dashLineRegex.MatchString(line) {
			continue
		}

		if patternLinksRegex.MatchString(line) {
			line = patternLinksRegex.ReplaceAllString(line, "")
		}

		if len(line) > 1 {
			if lastChar, ok := htGetLastChar(line); ok {
				if lastChar != '.' && lastChar != '?' && lastChar != ':' {
					line += "."
				} else if lastChar == ':' {
					line = line[0:len(line)-1] + "."
				}
			}
		}

		final += line + "\n"
	}

	ret := final
	ret = htReplaceRoman(ret)
	ret = htReplaceMarkdownLinks(ret)
	ret = strings.ReplaceAll(ret, "|", ".")
	ret = strings.ReplaceAll(ret, "*", "")
	ret = strings.ReplaceAll(ret, "( )", "")
	ret = htReplaceAllExceptions(ret, lang)
	ret = r.replaceMath(ret, lang)
	ret = htReplaceFractions(ret, r.dividedByKeyword(lang))
	ret = htConvertSuperscript(ret, lang)
	ret = htConvertFunctionAbbreviation(ret, lang)
	ret = htConvertTemperatures(ret, lang)
	ret = htRemoveHTMLTags(ret)
	ret = htReplaceSocialMediaUrls(ret)

	return ret
}

// dividedByKeyword returns the localized words used to read a fraction, for
// example "divided by" or "dividido por". The math keyword files are optional,
// so a language-aware default is used when they are not available.
func (r *runner) dividedByKeyword(lang string) string {
	_, _, _, dividedByStr := r.mathKeywords(lang)
	return strings.TrimSpace(dividedByStr)
}

func htGetLastChar(line string) (rune, bool) {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) == 0 {
		return 0, false
	}
	return []rune(trimmed)[len([]rune(trimmed))-1], true
}

var romanCache map[string]int = make(map[string]int)

var romanValues map[rune]int = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func htRomanToInt(roman string) int {
	if val, ok := romanCache[roman]; ok {
		return val
	}

	result := 0
	lastValue := 0
	for _, r := range roman {
		value := romanValues[r]
		if value > lastValue {
			result += value - 2*lastValue
		} else {
			result += value
		}
		lastValue = value
	}

	romanCache[roman] = result
	return result
}

func htReplaceRoman(text string) string {
	re := regexp.MustCompile(`\((Part|Parte) (I|II|III|IV|V|VI|VII|VIII|IX|X|XI|XII|XIII|XIV|XV|XVI|XVII|XVIII|XIX|XX)\)`)

	return re.ReplaceAllStringFunc(text, func(s string) string {
		submatches := re.FindStringSubmatch(s)
		if len(submatches) == 3 {
			part := submatches[1]
			roman := submatches[2]
			decimal := htRomanToInt(roman)
			return fmt.Sprintf("(%s %d)", part, decimal)
		}
		return s
	})
}

func htReplaceSocialMediaUrls(text string) string {
	socialPatterns := []struct {
		re   *regexp.Regexp
		name string
	}{
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?(?:bsky\.app|bluesky\.social)/\S*`), "Blue Sky"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?patreon\.com/\S*`), "Patreon"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?paypal\.com/\S*`), "PayPal"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?(?:creators\.)?spotify\.com/\S*`), "Spotify"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?facebook\.com/\S*`), "Facebook"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?twitter\.com/\S*`), "Twitter"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?x\.com/\S*`), "X"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?instagram\.com/\S*`), "Instagram"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?youtube\.com/\S*`), "You Tube"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?threads\.net/\S*`), "Threads"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?github\.com/\S*`), "GitHub"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?linkedin\.com/\S*`), "LinkedIn"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?tiktok\.com/\S*`), "TikTok"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?pinterest\.com/\S*`), "Pinterest"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?t\.me/\S*`), "Telegram"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?whatsapp\.com/\S*`), "WhatsApp"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?reddit\.com/\S*`), "Reddit"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?discord\.(?:com|gg)/\S*`), "Discord"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?telegram\.(?:me|org)/\S*`), "Telegram"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?signal\.org/\S*`), "Signal"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?mastodon\.\w+/\S*`), "Mastodon"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?snapchat\.com/\S*`), "Snapchat"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?tumblr\.com/\S*`), "Tumblr"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?medium\.com/\S*`), "Medium"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?bitbucket\.org/\S*`), "Bitbucket"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?gitlab\.com/\S*`), "GitLab"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?wikipedia\.org/\S*`), "Wikipedia"},
		{regexp.MustCompile(`\(?https?://(?:[a-z]+\.)?archive\.org/\S*`), "Archive"},
	}
	for _, sp := range socialPatterns {
		text = sp.re.ReplaceAllString(text, sp.name)
	}
	return text
}

func htReplaceMultiplication(text string, timesStr string) string {
	operand := `(?:` +
		`\b\d+(?:[.,/]\d+)*(?:[¹²³⁴⁵⁶⁷⁸⁹⁰]+)?%?` +
		`|\b[A-Za-z]\b` +
		`|\b[A-Za-z]+[ \t]*\([^()\n]*\)` +
		`|\([^()\n]*\)` +
		`)`

	// The letter x is only a multiplication signal when it is surrounded by
	// mathematical operands. This prevents ordinary words such as "xadrez"
	// or "aux" from being changed.
	xPattern := regexp.MustCompile(`(` + operand + `)[ \t]*\bx\b[ \t]*(` + operand + `)`)
	for {
		replaced := xPattern.ReplaceAllStringFunc(text, func(match string) string {
			parts := xPattern.FindStringSubmatch(match)
			return parts[1] + timesStr + parts[2]
		})
		if replaced == text {
			break
		}
		text = replaced
	}

	// An x without surrounding spaces ("5x6") is only valid between numbers;
	// the word-boundary form above cannot match it because digits are word
	// characters as well.
	compactXPattern := regexp.MustCompile(`(\b\d+(?:[.,/]\d+)*)[ \t]*x[ \t]*(\d+)`)
	for {
		replaced := compactXPattern.ReplaceAllStringFunc(text, func(match string) string {
			parts := compactXPattern.FindStringSubmatch(match)
			return parts[1] + timesStr + parts[2]
		})
		if replaced == text {
			break
		}
		text = replaced
	}

	// The multiplication sign is unambiguous, so it only needs operands
	// touching it on both sides. Prose mentions such as "the sign ×, which"
	// are left untouched because punctuation does not count as an operand.
	crossPattern := regexp.MustCompile(`([\p{L}\p{N})])[ \t]*×[ \t]*([\p{L}\p{N}(])`)
	for {
		replaced := crossPattern.ReplaceAllStringFunc(text, func(match string) string {
			parts := crossPattern.FindStringSubmatch(match)
			return parts[1] + timesStr + parts[2]
		})
		if replaced == text {
			break
		}
		text = replaced
	}

	return text
}

var mathOnlyLinePattern = regexp.MustCompile(`^[0-9\s.,+\-−×÷=/().\[\]%:²³⁴⁵⁶⁷⁸⁹⁰]+$`)

func htIsMathOnlyLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return false
	}

	return mathOnlyLinePattern.MatchString(trimmed) && strings.ContainsAny(trimmed, "0123456789")
}

func htReplaceSubtraction(text string, lang string) string {
	minusStr := " minus "
	if lang == "pt-BR" || lang == "es-ES" {
		minusStr = " menos "
	}

	operand := `(?:` +
		`\b\d+(?:[.,/]\d+)*(?:[¹²³⁴⁵⁶⁷⁸⁹⁰]+)?%?` +
		`|\b[A-Za-z]\b` +
		`|\b[A-Za-z]+[ \t]*\([^()\n]*\)` +
		`|\([^()\n]*\)` +
		`)`

	// An ASCII hyphen is also used for ranges and hyphenated words, so it is
	// only converted when the line is an equation (or a standalone math line
	// next to one). The Unicode minus sign is unambiguous and only needs
	// mathematical operands on both sides.
	hyphenPattern := regexp.MustCompile(`(` + operand + `)[ \t\x{00A0}]+-[ \t\x{00A0}]+(` + operand + `)`)
	minusPattern := regexp.MustCompile(`(` + operand + `)[ \t\x{00A0}]*−[ \t\x{00A0}]*` + `(` + operand + `)`)
	// In an equation, the Unicode minus can also separate word operands.
	equationMinusPattern := regexp.MustCompile(`([\p{L}\p{N})])[ \t\x{00A0}]*−[ \t\x{00A0}]*([\p{L}\p{N}(])`)

	lines := strings.Split(text, "\n")
	equationLine := make([]bool, len(lines))
	mathLine := make([]bool, len(lines))
	for i, line := range lines {
		equationLine[i] = strings.Contains(line, "=") ||
			strings.Contains(line, "+") ||
			strings.ContainsAny(line, "×÷")
		mathLine[i] = htIsMathOnlyLine(line)
	}

	neighborIsEquation := func(index int) bool {
		for i := index - 1; i >= 0; i-- {
			if strings.TrimSpace(lines[i]) == "" {
				continue
			}
			return equationLine[i]
		}
		return false
	}
	nextNeighborIsEquation := func(index int) bool {
		for i := index + 1; i < len(lines); i++ {
			if strings.TrimSpace(lines[i]) == "" {
				continue
			}
			return equationLine[i]
		}
		return false
	}

	replaceWith := func(line string, pattern *regexp.Regexp) string {
		for {
			replaced := pattern.ReplaceAllStringFunc(line, func(match string) string {
				parts := pattern.FindStringSubmatch(match)
				return parts[1] + minusStr + parts[2]
			})
			if replaced == line {
				return line
			}
			line = replaced
		}
	}

	for i := range lines {
		allowHyphen := equationLine[i] || (mathLine[i] && (neighborIsEquation(i) || nextNeighborIsEquation(i)))
		if allowHyphen {
			lines[i] = replaceWith(lines[i], hyphenPattern)
		}
		lines[i] = replaceWith(lines[i], minusPattern)
		if equationLine[i] {
			lines[i] = replaceWith(lines[i], equationMinusPattern)
		}
	}

	return strings.Join(lines, "\n")
}

// mathKeywordFallbacks returns the localized words used to read the four math
// operators supported by the audio generator (times, plus, equals, divided by)
// for when the optional math keyword files are not available.
func mathKeywordFallbacks(lang string) [4]string {
	switch lang {
	case "pt-BR":
		return [4]string{" vezes ", " mais ", " igual ", "dividido por"}
	case "es-ES":
		return [4]string{" veces ", " más ", " igual ", "dividido por"}
	default:
		return [4]string{" times ", " plus ", " equals ", "divided by"}
	}
}

// mathKeywords returns the localized math words. The values from the keyword
// file take precedence; missing values keep the language-aware fallback.
func (r *runner) mathKeywords(lang string) (string, string, string, string) {
	values := mathKeywordFallbacks(lang)

	localKeywords, err := r.loadKeywordsDirect(lang, "math_keywords")
	if err != nil {
		return values[0], values[1], values[2], values[3]
	}

	for i, index := range []int{33, 34, 35, 36} {
		if len(localKeywords) > index && strings.TrimSpace(localKeywords[index]) != "" {
			values[i] = localKeywords[index]
		}
	}

	return values[0], values[1], values[2], values[3]
}

func (r *runner) replaceMath(text string, lang string) string {
	timesStr, plusStr, equalsStr, _ := r.mathKeywords(lang)

	ret := htReplaceSubtraction(text, lang)
	ret = htReplaceMultiplication(ret, timesStr)

	ret = strings.ReplaceAll(ret, " + ", plusStr)
	ret = strings.ReplaceAll(ret, " = ", equalsStr)

	ret = strings.ReplaceAll(ret, "  +  ", " "+plusStr+" ")
	ret = strings.ReplaceAll(ret, "  =  ", " "+equalsStr+" ")

	ret = strings.ReplaceAll(ret, " +", " "+plusStr)
	ret = strings.ReplaceAll(ret, "+ ", plusStr+" ")

	ret = strings.ReplaceAll(ret, " =", " "+equalsStr)
	ret = strings.ReplaceAll(ret, "= ", equalsStr+" ")

	return ret
}

// htIsInternalContentLink reports whether href points to the History Tracers
// index. Those links are navigation aids for the web page and should not be
// read aloud in the generated audio.
func htIsInternalContentLink(href string) bool {
	trimmed := strings.TrimSpace(href)
	if trimmed == "" {
		return false
	}

	parsed, err := url.Parse(trimmed)
	if err != nil {
		return false
	}

	if parsed.Host != "" {
		host := strings.ToLower(parsed.Hostname())
		if host != "historytracers.org" && host != "www.historytracers.org" {
			return false
		}
	}

	path := strings.ToLower(parsed.Path)
	for strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../") {
		path = strings.TrimPrefix(path, "./")
		path = strings.TrimPrefix(path, "../")
	}
	path = strings.TrimPrefix(path, "/")

	return path == "index.html" || strings.HasSuffix(path, "/index.html")
}

var markdownLinkPattern = regexp.MustCompile(`!?\[([^\]]*)\]\(\s*([^)\s]+)(?:\s+"[^"]*")?\s*\)`)

// htReplaceMarkdownLinks removes the URL of internal navigation links, so the
// anchor text is read aloud without the index.html address. Other markdown
// links are left untouched.
func htReplaceMarkdownLinks(text string) string {
	return markdownLinkPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := markdownLinkPattern.FindStringSubmatch(match)
		if len(parts) != 3 || strings.HasPrefix(match, "!") {
			return match
		}
		if htIsInternalContentLink(parts[2]) {
			return parts[1]
		}
		return match
	})
}

func htIsFullyParenthesized(text string) bool {
	text = strings.TrimSpace(text)
	if len(text) < 2 || text[0] != '(' || text[len(text)-1] != ')' {
		return false
	}

	depth := 0
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 && i != len(text)-1 {
				return false
			}
		}
	}

	return depth == 0
}

func htRenderMathFraction(numerator string, denominator string, dividedByStr string) string {
	needsParentheses := func(expression string) bool {
		if strings.Contains(expression, dividedByStr) {
			return true
		}
		return strings.ContainsAny(expression, "+−-×÷=")
	}

	if needsParentheses(numerator) && !htIsFullyParenthesized(numerator) {
		numerator = "(" + numerator + ")"
	}
	if needsParentheses(denominator) && !htIsFullyParenthesized(denominator) {
		denominator = "(" + denominator + ")"
	}

	return numerator + " " + dividedByStr + " " + denominator
}

var plainFractionPattern = regexp.MustCompile(`(^|[\s(])(\d+)\s*(?:/|⁄|&frasl;|&#8260;)\s*(\d+)([)\s.,;:!?]|$)`)

// htReplaceFractions turns plain fractions such as "1/2" into the spoken form
// "1 divided by 2". Markdown renders those fractions as superscript and
// subscript tags on the website; the smartphone text still contains the
// original slash, so both spellings are handled here.
func htReplaceFractions(text string, dividedByStr string) string {
	return plainFractionPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := plainFractionPattern.FindStringSubmatch(match)
		if len(parts) != 5 {
			return match
		}
		return parts[1] + htRenderMathFraction(parts[2], parts[3], dividedByStr) + parts[4]
	})
}

// loadKeywordsDirect reads the common/math keyword file of a language. The
// keyword files are optional: when they are not present a nil slice and an
// error are returned, and the caller falls back to its defaults.
func (r *runner) loadKeywordsDirect(lang string, name string) ([]string, error) {
	fileName := fmt.Sprintf("%slang/%s/%s.json", r.cfg.SrcPath, lang, name)

	byteValue, err := htOpenFileReadClose(fileName)
	if err != nil {
		return nil, err
	}

	var keywords []string
	if name == "math_keywords" {
		var kf common.HTMathKeywordsFormat
		if err := json.Unmarshal(byteValue, &kf); err != nil {
			htCommonJSONError(byteValue, err)
			return nil, err
		}
		keywords = kf.Keywords
	} else {
		var kf common.HTKeywordsFormat
		if err := json.Unmarshal(byteValue, &kf); err != nil {
			htCommonJSONError(byteValue, err)
			return nil, err
		}
		keywords = kf.Keywords
	}

	return keywords, nil
}

func htReplaceAllExceptions(text string, lang string) string {
	ret := strings.ReplaceAll(text, "(#)", "")

	for _, letter := range []string{"α", "β", "γ", "δ", "ε", "ζ", "η", "θ", "ι", "κ", "λ", "μ", "ν", "ξ", "π", "ρ", "σ", "τ", "υ", "φ", "χ", "ψ", "ω"} {
		ret = strings.ReplaceAll(ret, letter, htConvertGreekLetter(letter, lang))
	}

	for _, mathSymbol := range []string{"ℕ", "ℤ", "ℚ", "ℝ", "ℂ", "ℙ", "ℕ₀", "ℕ*", "ℕ⁺"} {
		ret = strings.ReplaceAll(ret, mathSymbol, "")
	}

	ret = htRemoveDuplicateParentheses(ret)

	return ret
}

func htRemoveDuplicateParentheses(text string) string {
	duplicateParenRegex := regexp.MustCompile(`(\w+)\s*\(\s*(\w+)\s*\)`)
	return duplicateParenRegex.ReplaceAllStringFunc(text, func(match string) string {
		matches := duplicateParenRegex.FindStringSubmatch(match)
		if len(matches) == 3 && matches[1] == matches[2] {
			return matches[1]
		}
		return match
	})
}

func htConvertGreekLetter(letter string, lang string) string {
	greekLetterMap := map[string]map[string]string{
		"α": {"en-US": "alpha", "es-ES": "alfa", "pt-BR": "alfa"},
		"β": {"en-US": "beta", "es-ES": "beta", "pt-BR": "beta"},
		"γ": {"en-US": "gamma", "es-ES": "gama", "pt-BR": "gama"},
		"δ": {"en-US": "delta", "es-ES": "delta", "pt-BR": "delta"},
		"ε": {"en-US": "epsilon", "es-ES": "épsilon", "pt-BR": "epsilon"},
		"ζ": {"en-US": "zeta", "es-ES": "zeta", "pt-BR": "zeta"},
		"η": {"en-US": "eta", "es-ES": "eta", "pt-BR": "eta"},
		"θ": {"en-US": "theta", "es-ES": "theta", "pt-BR": "teta"},
		"ι": {"en-US": "iota", "es-ES": "iota", "pt-BR": "iota"},
		"κ": {"en-US": "kappa", "es-ES": "kappa", "pt-BR": "kappa"},
		"λ": {"en-US": "lambda", "es-ES": "lambda", "pt-BR": "lambda"},
		"μ": {"en-US": "mu", "es-ES": "mu", "pt-BR": "mi"},
		"ν": {"en-US": "nu", "es-ES": "nu", "pt-BR": "ni"},
		"ξ": {"en-US": "xi", "es-ES": "xi", "pt-BR": "xi"},
		"π": {"en-US": "pi", "es-ES": "pi", "pt-BR": "pi"},
		"ρ": {"en-US": "rho", "es-ES": "rho", "pt-BR": "ro"},
		"σ": {"en-US": "sigma", "es-ES": "sigma", "pt-BR": "sigma"},
		"τ": {"en-US": "tau", "es-ES": "tau", "pt-BR": "tau"},
		"υ": {"en-US": "upsilon", "es-ES": "ípsilon", "pt-BR": "ípsilon"},
		"φ": {"en-US": "phi", "es-ES": "fi", "pt-BR": "fi"},
		"ϕ": {"en-US": "phi", "es-ES": "fi", "pt-BR": "fi"},
		"χ": {"en-US": "chi", "es-ES": "ji", "pt-BR": "qui"},
		"ψ": {"en-US": "psi", "es-ES": "psi", "pt-BR": "psi"},
		"ω": {"en-US": "omega", "es-ES": "omega", "pt-BR": "omega"},
	}

	if langMap, ok := greekLetterMap[letter]; ok {
		if name, ok := langMap[lang]; ok {
			return name
		}
	}
	return letter
}

func htConvertSuperscript(text string, lang string) string {
	superscriptMap := map[string]map[string]string{
		"¹":      {"en-US": " to the first power", "es-ES": " a la primera potencia", "pt-BR": " à primeira potência"},
		"²":      {"en-US": " squared", "es-ES": " al cuadrado", "pt-BR": " ao quadrado"},
		"³":      {"en-US": " cubed", "es-ES": " al cubo", "pt-BR": " ao cubo"},
		"⁴":      {"en-US": " to the fourth power", "es-ES": " a la cuarta potencia", "pt-BR": " à quarta potência"},
		"⁵":      {"en-US": " to the fifth power", "es-ES": " a la quinta potencia", "pt-BR": " à quinta potência"},
		"⁶":      {"en-US": " to the sixth power", "es-ES": " a la sexta potencia", "pt-BR": " à sexta potência"},
		"⁷":      {"en-US": " to the seventh power", "es-ES": " a la septima potencia", "pt-BR": " à sétima potência"},
		"⁸":      {"en-US": " to the eighth power", "es-ES": " a la octava potencia", "pt-BR": " à oitava potência"},
		"⁹":      {"en-US": " to the ninth power", "es-ES": " a la novena potencia", "pt-BR": " à nona potência"},
		"⁰":      {"en-US": " to the zero power", "es-ES": " a la cero potencia", "pt-BR": " à zero potência"},
		"&sup1;": {"en-US": " to the first power", "es-ES": " a la primera potencia", "pt-BR": " à primeira potência"},
		"&sup2;": {"en-US": " squared", "es-ES": " al cuadrado", "pt-BR": " ao quadrado"},
		"&sup3;": {"en-US": " cubed", "es-ES": " al cubo", "pt-BR": " ao cubo"},
	}

	implicitMultPattern := regexp.MustCompile(`(\w+)([¹²³⁴⁵⁶⁷⁸⁹⁰])\s+(\w+)([¹²³⁴⁵⁶⁷⁸⁹⁰])\s*\(`)

	timesStr := " vezes "
	if lang == "en-US" {
		timesStr = " times "
	} else if lang == "es-ES" {
		timesStr = " por "
	}

	text = implicitMultPattern.ReplaceAllString(text, "$1$2 "+timesStr+"$3$4(")

	reversePattern := regexp.MustCompile(`(\w+)([¹²³⁴⁵⁶⁷⁸⁹⁰])\s*\(\s*([^)]+)\s*\)\s+(\w+)([¹²³⁴⁵⁶⁷⁸⁹⁰])`)
	text = reversePattern.ReplaceAllString(text, "$1$2($3) "+timesStr+"$4$5")

	for superscript, langMap := range superscriptMap {
		if replacement, ok := langMap[lang]; ok {
			text = strings.ReplaceAll(text, superscript, replacement)
		}
	}

	return text
}

func htConvertFunctionAbbreviation(text string, lang string) string {
	preposition := "de"
	if lang == "en-US" {
		preposition = "of"
	}

	powerMap := map[string]map[string]string{
		"ao quadrado": {"en-US": "squared", "es-ES": "al cuadrado", "pt-BR": "ao quadrado"},
		"ao cubo":     {"en-US": "cubed", "es-ES": "al cubo", "pt-BR": "ao cubo"},
		"squared":     {"en-US": "squared", "es-ES": "al cuadrado", "pt-BR": "ao quadrado"},
		"cubed":       {"en-US": "cubed", "es-ES": "al cubo", "pt-BR": "ao cubo"},
	}

	funcMap := map[string]map[string]string{
		"cos":    {"en-US": "cosine", "es-ES": "coseno", "pt-BR": "cosseno"},
		"sin":    {"en-US": "sine", "es-ES": "seno", "pt-BR": "seno"},
		"sen":    {"en-US": "sine", "es-ES": "seno", "pt-BR": "seno"},
		"tan":    {"en-US": "tangent", "es-ES": "tangente", "pt-BR": "tangente"},
		"sec":    {"en-US": "secant", "es-ES": "secante", "pt-BR": "secante"},
		"csc":    {"en-US": "cosecant", "es-ES": "cosecante", "pt-BR": "cossecante"},
		"cot":    {"en-US": "cotangent", "es-ES": "cotangente", "pt-BR": "cotangente"},
		"arccos": {"en-US": "arc cosine", "es-ES": "arcocoseno", "pt-BR": "arcocosseno"},
		"arcsin": {"en-US": "arc sine", "es-ES": "arcoseno", "pt-BR": "arcosseno"},
		"arctan": {"en-US": "arc tangent", "es-ES": "arcotangente", "pt-BR": "arcotangente"},
		"log":    {"en-US": "logarithm", "es-ES": "logaritmo", "pt-BR": "logaritmo"},
		"ln":     {"en-US": "natural logarithm", "es-ES": "logaritmo natural", "pt-BR": "logaritmo natural"},
		"exp":    {"en-US": "exponential", "es-ES": "exponencial", "pt-BR": "exponencial"},
		"sqrt":   {"en-US": "square root", "es-ES": "raíz cuadrada", "pt-BR": "raiz quadrada"},
		"abs":    {"en-US": "absolute value", "es-ES": "valor absoluto", "pt-BR": "valor absoluto"},
		"max":    {"en-US": "maximum", "es-ES": "máximo", "pt-BR": "máximo"},
		"min":    {"en-US": "minimum", "es-ES": "mínimo", "pt-BR": "mínimo"},
		"mod":    {"en-US": "modulo", "es-ES": "módulo", "pt-BR": "módulo"},
	}

	for abbr, langMap := range funcMap {
		if full, ok := langMap[lang]; ok {
			pattern := regexp.MustCompile(`\b` + abbr + `\s*\(\s*([^)]+)\s*\)`)
			text = pattern.ReplaceAllString(text, full+" "+preposition+" $1")

			superscriptPattern := regexp.MustCompile(`\b` + abbr + `([¹²³⁴⁵⁶⁷⁸⁹⁰])\s*\(\s*([^)]+)\s*\)`)
			text = superscriptPattern.ReplaceAllString(text, full+"$1 "+preposition+" $2")
		}
	}

	for power, powerLangMap := range powerMap {
		if powerWord, ok := powerLangMap[lang]; ok {
			powerPattern := regexp.MustCompile(`(\w+)\s+` + power + `\s*\(\s*([^)]+)\s*\)`)
			text = powerPattern.ReplaceAllString(text, "$1 "+powerWord+" "+preposition+" $2")
		}
	}

	return text
}

func htConvertTemperatures(text string, lang string) string {
	tempRegex := regexp.MustCompile(`(-?\d+(?:\.\d+)?)\s*º?\s*([CF])\s*\(\s*(-?\d+(?:\.\d+)?)\s*º?\s*([CF])\)`)

	tempMap := map[string]map[string]string{
		"celsius": {
			"pt-BR": "graus celsius",
			"es-ES": "grados celsius",
			"en-US": "degrees celsius",
		},
		"fahrenheit": {
			"pt-BR": "graus fahrenheit",
			"es-ES": "grados fahrenheit",
			"en-US": "degrees fahrenheit",
		},
		"negativos": {
			"pt-BR": "negativos",
			"es-ES": "negativos",
			"en-US": "negative",
		},
		"positivos": {
			"pt-BR": "positivos",
			"es-ES": "positivos",
			"en-US": "positive",
		},
		"correspond": {
			"pt-BR": "que correspondem a",
			"es-ES": "que corresponden a",
			"en-US": "which correspond to",
		},
	}

	return tempRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := tempRegex.FindStringSubmatch(match)
		if len(parts) != 5 {
			return match
		}

		temp1Str := parts[1]
		unit1 := strings.ToUpper(parts[2])
		temp2Str := parts[3]

		temp1, _ := strconv.ParseFloat(temp1Str, 64)

		var unit1Name, unit2Name, negPos string

		if unit1 == "C" {
			unit1Name = tempMap["celsius"][lang]
			unit2Name = tempMap["fahrenheit"][lang]
		} else {
			unit1Name = tempMap["fahrenheit"][lang]
			unit2Name = tempMap["celsius"][lang]
		}

		if temp1 < 0 {
			negPos = tempMap["negativos"][lang]
		} else {
			negPos = tempMap["positivos"][lang]
		}

		corrText := tempMap["correspond"][lang]

		absTemp1 := int(math.Abs(temp1))
		absTemp2, _ := strconv.Atoi(temp2Str)

		return fmt.Sprintf("%d %s %s (%s %d %s %s)", absTemp1, unit1Name, negPos, corrText, absTemp2, unit2Name, negPos)
	})
}

func htRemoveAsiaticCharacters(text string) string {
	text = strings.ReplaceAll(text, "Schyoty (счёты)", "Schyoty")
	text = strings.ReplaceAll(text, "(матрёшка)", "")

	chineseRegex := regexp.MustCompile(`[\p{Han}]+`)
	cleaned := chineseRegex.ReplaceAllString(text, "")

	emojiSymbolRegex := regexp.MustCompile(`[\x{1F000}-\x{1FAFF}\x{2600}-\x{27BF}\x{2B00}-\x{2BFF}\x{FE0F}\x{200D}]`)
	cleaned = emojiSymbolRegex.ReplaceAllString(cleaned, "")

	emptyParenRegex := regexp.MustCompile(`\(\s*\)`)
	cleaned = emptyParenRegex.ReplaceAllString(cleaned, "")

	return cleaned
}

func htRemoveSVGBlocks(text string) string {
	svgRegex := regexp.MustCompile(`(?s)<svg[^>]*>.*?</svg>`)
	return svgRegex.ReplaceAllString(text, "")
}

func htRemoveStyleBlocks(text string) string {
	styleRegex := regexp.MustCompile(`(?s)<style[^>]*>.*?</style>`)
	return styleRegex.ReplaceAllString(text, "")
}

func htRemoveHTMLTags(text string) string {
	text = htRemoveSVGBlocks(text)
	text = htRemoveStyleBlocks(text)
	// Only match actual HTML tags (an opening '<' optionally followed by '/',
	// then a tag name starting with a letter). Matching any '<...>' span would
	// swallow literal comparison symbols such as '(<)' or '0 < 1' together with
	// everything up to the next '>' in the surrounding plain text.
	htmlTagRegex := regexp.MustCompile(`</?[a-zA-Z][^>]*>`)
	return htmlTagRegex.ReplaceAllString(text, "")
}

func htAdjustTrailingDots(text string) string {
	lines := strings.Split(text, "\n")
	var result []string
	var prevWasEmpty bool

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "." {
			continue
		}

		dotRegex := regexp.MustCompile(`\.{2,}$`)
		if dotRegex.MatchString(line) {
			line = dotRegex.ReplaceAllString(line, ".")
		}

		if trimmed == "" {
			if !prevWasEmpty {
				result = append(result, line)
			}
			prevWasEmpty = true
		} else {
			result = append(result, line)
			prevWasEmpty = false
		}
	}

	return strings.Join(result, "\n")
}
