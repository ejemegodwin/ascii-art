package main

import (
	"os"
	"testing"
)

func TestLoadBanner_ValidFile(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(banner) != 95 {
		t.Errorf("expected 95 characters, got %d", len(banner))
	}
}

func TestLoadBanner_MissingFile(t *testing.T) {
	_, err := LoadBanner("nonexistent.txt")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	tmp, _ := os.CreateTemp("", "empty_banner_*.txt")
	defer os.Remove(tmp.Name())
	tmp.Close()

	_, err := LoadBanner(tmp.Name())
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}

func TestLoadBanner_SpaceCharExists(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if _, ok := banner[' ']; !ok {
		t.Error("expected space character to exist in banner")
	}
}

func TestLoadBanner_TildeCharExists(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if _, ok := banner['~']; !ok {
		t.Error("expected tilde character to exist in banner")
	}
}

func TestLoadBanner_EachCharHas8Lines(t *testing.T) {
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	for r := rune(32); r <= 126; r++ {
		art, ok := banner[r]
		if !ok {
			t.Errorf("missing character %q in banner", r)
			continue
		}
		if len(art) != 8 {
			t.Errorf("character %q has %d lines, expected 8", r, len(art))
		}
	}
}
