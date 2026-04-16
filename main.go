package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	// Case 1: With output flag
	if len(args) == 3 && strings.HasPrefix(args[0], "--output=") {
		outputFile := strings.TrimPrefix(args[0], "--output=")
		input := args[1]
		banner := args[2]

		if outputFile == "" {
			usage()
			return
		}

		result, err := generateAscii(input, banner)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = os.WriteFile(outputFile, []byte(result), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		return
	}

	// Case 2: Without flag
	if len(args) == 1 {
		result, err := generateAscii(args[0], "standard")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(result)
		return
	}

	// Invalid usage
	usage()
}

func usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
}

func generateAscii(input, banner string) (string, error) {
	font, err := loadBanner(banner + ".txt")
	if err != nil {
		return "", err
	}

	lines := strings.Split(input, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range line {
				if ch < 32 || ch > 126 {
					return "", fmt.Errorf("unsupported character: %q", ch)
				}
				index := (int(ch) - 32) * 9
				result.WriteString(font[index+i])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

func loadBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading banner file: %s", filename)
	}

	lines := strings.Split(string(data), "\n")

	// Each character = 8 lines + 1 empty line → total 9 lines per char
	if len(lines) < 95*9 {
		return nil, fmt.Errorf("invalid banner file")
	}

	return lines, nil
}
