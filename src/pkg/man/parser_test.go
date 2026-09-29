package man

import (
	"testing"
)

const sampleManPage = `LS(1)                            User Commands                           LS(1)

NAME
     ls - list directory contents

SYNOPSIS
     ls [OPTION]... [FILE]...

DESCRIPTION
     List information about the FILEs.

OPTIONS
     -a, --all
            do not ignore entries starting with .

     -l
            use a long listing format

     --block-size=SIZE
            scale sizes by SIZE when printing them

EXIT STATUS
     0      if OK,
     1      if minor problems.
`

func TestCleanFormatting(t *testing.T) {
	raw := "B\bBO\bOL\bLD\bD and \x1b[31mRED\x1b[0m"
	cleaned := CleanFormatting(raw)
	assertMatch(t, cleaned, "BOLD and RED")
}

func TestParseSections(t *testing.T) {
	sections := ParseSections(sampleManPage)
	assertContains(t, sections, "NAME")
	assertContains(t, sections, "DESCRIPTION")
	assertContains(t, sections, "OPTIONS")
}

func TestExtractSection(t *testing.T) {
	desc, ok := ExtractSection(sampleManPage, "description")
	assertTrue(t, ok)
	assertMatch(t, desc, "List information about the FILEs.")
}

func TestExtractFlag(t *testing.T) {
	flagA, ok := ExtractFlag(sampleManPage, "-a")
	assertTrue(t, ok)
	assertContainsString(t, flagA, "do not ignore entries")

	flagBlock, okBlock := ExtractFlag(sampleManPage, "block-size")
	assertTrue(t, okBlock)
	assertContainsString(t, flagBlock, "scale sizes by SIZE")

	_, notFound := ExtractFlag(sampleManPage, "-z")
	assertFalse(t, notFound)
}

func assertMatch(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func assertTrue(t *testing.T, condition bool) {
	if !condition {
		t.Fatal("expected true, got false")
	}
}

func assertFalse(t *testing.T, condition bool) {
	if condition {
		t.Fatal("expected false, got true")
	}
}

func assertContains(t *testing.T, m map[string]string, key string) {
	if _, ok := m[key]; !ok {
		t.Fatalf("expected key %q in map", key)
	}
}

func assertContainsString(t *testing.T, text, substr string) {
	if !containsSubstr(text, substr) {
		t.Fatalf("expected %q to contain %q", text, substr)
	}
}

func containsSubstr(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || stringSearch(s, sub))
}

func stringSearch(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
