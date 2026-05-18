package main

import "testing"

func TestSplitInput_SingleWord(t *testing.T) {
	parts := SplitInput("Hello")
	if len(parts) != 1 || parts[0] != "Hello" {
		t.Errorf("expected [Hello], got %v", parts)
	}
}

func TestSplitInput_WithNewline(t *testing.T) {
	parts := SplitInput(`Hello\nWorld`)
	if len(parts) != 2 {
		t.Fatalf("expected 2 parts, got %d", len(parts))
	}
	if parts[0] != "Hello" || parts[1] != "World" {
		t.Errorf("unexpected parts: %v", parts)
	}
}

func TestSplitInput_MultipleNewlines(t *testing.T) {
	parts := SplitInput(`A\n\nB`)
	if len(parts) != 3 {
		t.Fatalf("expected 3 parts, got %d", len(parts))
	}
	if parts[1] != "" {
		t.Errorf("expected middle part to be empty, got %q", parts[1])
	}
}

func TestSplitInput_EmptyString(t *testing.T) {
	parts := SplitInput("")
	if len(parts) != 1 || parts[0] != "" {
		t.Errorf("expected [\"\"], got %v", parts)
	}
}

func TestValidateInput_ValidASCII(t *testing.T) {
	_, err := ValidateInput("Hello, World! 123")
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestValidateInput_AllPrintable(t *testing.T) {
	all := ""
	for r := rune(32); r <= 126; r++ {
		all += string(r)
	}
	_, err := ValidateInput(all)
	if err != nil {
		t.Errorf("expected no error for all printable ASCII, got: %v", err)
	}
}

func TestValidateInput_InvalidCharacter(t *testing.T) {
	_, err := ValidateInput("Hello\x01World")
	if err == nil {
		t.Error("expected error for invalid character, got nil")
	}
}

func TestValidateInput_EmptyString(t *testing.T) {
	_, err := ValidateInput("")
	if err != nil {
		t.Errorf("expected no error for empty string, got: %v", err)
	}
}

func TestValidateInput_UnicodeCharacter(t *testing.T) {
	_, err := ValidateInput("héllo")
	if err == nil {
		t.Error("expected error for unicode character, got nil")
	}
}
