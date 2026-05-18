package main

import (
	"strings"
	"testing"
)

func TestFullPipeline_Hello(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	_, err = ValidateInput("Hello")
	if err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	result := GenerateArt("Hello", banner)
	if result == "" {
		t.Error("expected non-empty output for 'Hello'")
	}
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

func TestFullPipeline_WithNewline(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("failed to load banner: %v", err)
	}

	input := `Hello\nThere`
	_, err = ValidateInput(input)
	if err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	result := GenerateArt(input, banner)
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(lines) != 16 {
		t.Errorf("expected 16 lines for two-word input, got %d", len(lines))
	}
}

func TestFullPipeline_InvalidInputRejected(t *testing.T) {
	_, err := ValidateInput("Hello\x80World")
	if err == nil {
		t.Error("expected validation to reject invalid character")
	}
}

func TestFullPipeline_ShadowBanner(t *testing.T) {
	banner, err := LoadBanner("shadow.txt")
	if err != nil {
		t.Fatalf("failed to load shadow banner: %v", err)
	}
	result := GenerateArt("Hi", banner)
	if result == "" {
		t.Error("expected non-empty output from shadow banner")
	}
}

func TestFullPipeline_ThinkertoyBanner(t *testing.T) {
	banner, err := LoadBanner("thinkertoy.txt")
	if err != nil {
		t.Fatalf("failed to load thinkertoy banner: %v", err)
	}
	result := GenerateArt("Hi", banner)
	if result == "" {
		t.Error("expected non-empty output from thinkertoy banner")
	}
}
