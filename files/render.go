package main

func RenderLine(text string, banner map[rune][]string) []string {
	lines := make([]string, 8)
	for _, r := range text {
		art, ok := banner[r]
		if !ok {
			continue
		}
		for i := 0; i < 8; i++ {
			lines[i] += art[i]
		}
	}
	return lines
}
