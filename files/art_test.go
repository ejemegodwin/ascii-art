package main

import (
	"strings"
	"testing"
)

func setupArtBanner(t *testing.T) map[rune][]string {
	t.Helper()
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}
	return banner
}

func TestGenerateArt_EmptyInput(t *testing.T) {
	banner := setupArtBanner(t)
	result := GenerateArt("", banner)
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestGenerateArt_SingleWord_Has8Lines(t *testing.T) {
	banner := setupArtBanner(t)
	result := GenerateArt("Hi", banner)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

func TestGenerateArt_NewlineSplitsIntoTwoBlocks(t *testing.T) {
	banner := setupArtBanner(t)
	result := GenerateArt(`Hello\nWorld`, banner)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	// 8 lines for Hello + 8 lines for World = 16
	if len(lines) != 16 {
		t.Errorf("expected 16 lines, got %d: %v", len(lines), lines)
	}
}

func TestGenerateArt_DoubleNewlineAddsBlankLine(t *testing.T) {
	banner := setupArtBanner(t)
	result := GenerateArt(`Hello\n\nWorld`, banner)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	// 8 lines for Hello + 1 blank + 8 lines for World = 17
	if len(lines) != 17 {
		t.Errorf("expected 17 lines, got %d", len(lines))
	}
}

func TestGenerateArt_OnlyNewline(t *testing.T) {
	banner := setupArtBanner(t)
	result := GenerateArt(`\n`, banner)
	if result != "\n" {
		t.Errorf("expected single newline, got %q", result)
	}
}

func TestGenerateArt_OutputEndsWithNewline(t *testing.T) {
	banner := setupArtBanner(t)
	result := GenerateArt("A", banner)
	if !strings.HasSuffix(result, "\n") {
		t.Errorf("expected output to end with newline, got %q", result)
	}
}