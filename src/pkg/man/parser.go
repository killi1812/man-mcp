package man

import (
	"regexp"
	"strings"
)

var (
	ansiRegex    = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	sectionRegex = regexp.MustCompile(`^[A-Z][A-Z0-9 _-]{1,30}$`)
	optRegex     = regexp.MustCompile(`^\s*(-[a-zA-Z0-9]|--[a-zA-Z0-9_-]+)`)
)

// CleanFormatting strips ANSI codes and terminal overstrikes.
func CleanFormatting(input string) string {
	cleaned := stripOverstrike(input)
	return stripANSI(cleaned)
}

func stripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

func stripOverstrike(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	sb.Grow(len(runes))
	processRunes(runes, &sb)
	return sb.String()
}

func processRunes(runes []rune, sb *strings.Builder) {
	for i := 0; i < len(runes); i++ {
		i = handleRuneAt(runes, i, sb)
	}
}

func handleRuneAt(runes []rune, i int, sb *strings.Builder) int {
	if isOverstrikeSeq(runes, i) {
		return i + 1
	}
	sb.WriteRune(runes[i])
	return i
}

func isOverstrikeSeq(runes []rune, i int) bool {
	return i+1 < len(runes) && runes[i+1] == '\b'
}

// ParseSections extracts major sections into a map.
func ParseSections(content string) map[string]string {
	clean := CleanFormatting(content)
	lines := strings.Split(clean, "\n")
	sections := make(map[string]string)
	collectSections(lines, sections)
	return sections
}

func collectSections(lines []string, sections map[string]string) {
	currentHeader := ""
	var currentLines []string
	for _, line := range lines {
		currentHeader, currentLines = processSectionLine(line, currentHeader, currentLines, sections)
	}
	saveSection(currentHeader, currentLines, sections)
}

func processSectionLine(line string, header string, lines []string, m map[string]string) (string, []string) {
	if isSectionHeader(line) {
		saveSection(header, lines, m)
		return strings.TrimSpace(line), nil
	}
	return header, append(lines, line)
}

func isSectionHeader(line string) bool {
	trimmed := strings.TrimRight(line, " \t\r")
	if len(trimmed) == 0 || line[0] == ' ' || line[0] == '\t' {
		return false
	}
	return sectionRegex.MatchString(trimmed)
}

func saveSection(header string, lines []string, m map[string]string) {
	if header != "" && len(lines) > 0 {
		m[header] = strings.TrimSpace(strings.Join(lines, "\n"))
	}
}

// ExtractSection retrieves the text of a named section case-insensitively.
func ExtractSection(content string, sectionName string) (string, bool) {
	sections := ParseSections(content)
	target := strings.ToUpper(strings.TrimSpace(sectionName))
	val, ok := sections[target]
	return val, ok
}

// ExtractFlag searches for a specific flag inside options or content.
func ExtractFlag(content string, flag string) (string, bool) {
	clean := CleanFormatting(content)
	normFlag := normalizeFlag(flag)
	lines := strings.Split(clean, "\n")
	return findFlagBlock(lines, normFlag)
}

func normalizeFlag(flag string) string {
	if !strings.HasPrefix(flag, "-") {
		return "--" + flag
	}
	return flag
}

func findFlagBlock(lines []string, flag string) (string, bool) {
	startIdx := locateFlagLine(lines, flag)
	if startIdx == -1 {
		return "", false
	}
	return extractBlockFrom(lines, startIdx), true
}

func locateFlagLine(lines []string, flag string) int {
	for i, line := range lines {
		if lineMatchesFlag(line, flag) {
			return i
		}
	}
	return -1
}

func lineMatchesFlag(line string, flag string) bool {
	trimmed := strings.TrimSpace(line)
	if !optRegex.MatchString(line) {
		return false
	}
	return flagInLine(trimmed, flag)
}

func flagInLine(line string, flag string) bool {
	words := strings.Fields(line)
	for _, word := range words {
		if cleanFlagWord(word) == flag {
			return true
		}
	}
	return false
}

func cleanFlagWord(word string) string {
	base := cutFlagWord(word)
	return strings.Trim(base, ", \t")
}

func cutFlagWord(word string) string {
	if idx := strings.IndexAny(word, "=:["); idx != -1 {
		return word[:idx]
	}
	return word
}


func extractBlockFrom(lines []string, start int) string {
	var block []string
	block = append(block, lines[start])
	for i := start + 1; i < len(lines); i++ {
		if shouldStopFlagBlock(lines[i]) {
			break
		}
		block = append(block, lines[i])
	}
	return strings.TrimSpace(strings.Join(block, "\n"))
}

func shouldStopFlagBlock(line string) bool {
	if isSectionHeader(line) {
		return true
	}
	return optRegex.MatchString(line)
}
