package main

import "testing"

func setupBanner(t *testing.T) map[rune][]string {
	t.Helper()
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}
	return banner
}

func TestRenderLine_Returns8Lines(t *testing.T) {
	banner := setupBanner(t)
	lines := RenderLine("A", banner)
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

func TestRenderLine_EmptyString(t *testing.T) {
	banner := setupBanner(t)
	lines := RenderLine("", banner)
	for i, line := range lines {
		if line != "" {
			t.Errorf("expected empty line at index %d, got %q", i, line)
		}
	}
}

func TestRenderLine_SpaceCharacter(t *testing.T) {
	banner := setupBanner(t)
	lines := RenderLine(" ", banner)
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

func TestRenderLine_MultipleCharsWidthGrows(t *testing.T) {
	banner := setupBanner(t)
	one := RenderLine("A", banner)
	two := RenderLine("AB", banner)
	if len(two[0]) <= len(one[0]) {
		t.Errorf("expected wider output for 2 chars, one=%d two=%d", len(one[0]), len(two[0]))
	}
}

func TestRenderLine_UnknownCharSkipped(t *testing.T) {
	banner := setupBanner(t)
	linesWithout := RenderLine("AB", banner)
	// inserting a rune outside range should be skipped, width unchanged
	delete(banner, 'B')
	linesMissing := RenderLine("AB", banner)
	if len(linesMissing[0]) >= len(linesWithout[0]) {
		t.Errorf("expected narrower output when character is missing from banner")
	}
}
