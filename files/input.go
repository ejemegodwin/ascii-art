package main

import (
	"fmt"
	"strings"
)

func SplitInput(input string) []string {
	return strings.Split(input, `\n`)
}

func ValidateInput(s string) (rune, error) {
	for _, r := range s {
		if r == '\n' {
			continue
		}
		if r < 32 || r > 126 {
			return r, fmt.Errorf("unsupported character: %c", r)
		}
	}
	return 0, nil
}
