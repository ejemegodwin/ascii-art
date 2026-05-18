package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	parts := SplitInput(input)
	var out strings.Builder

	for i, part := range parts {
		if part == "" {
			if i < len(parts)-1 {
				out.WriteString("\n")
			}
			continue
		}
		rendered := RenderLine(part, banner)
		for _, line := range rendered {
			out.WriteString(line)
			out.WriteString("\n")
		}
	}

	return out.String()
}
