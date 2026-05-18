package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading %q banner", filename)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty banner file")
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	banner := make(map[rune][]string)
	for ascii := 32; ascii <= 126; ascii++ {
		start := (ascii-32)*9 + 1
		end := start + 8
		if end > len(lines) {
			return nil, fmt.Errorf("invalid banner content")
		}
		banner[rune(ascii)] = lines[start:end]
	}
	return banner, nil
}
