package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run . [STRING] [BANNER]")
		os.Exit(1)
	}

	input := os.Args[1]

	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		bannerFile = os.Args[2] + ".txt"
	}

	if _, err := ValidateInput(input); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	banner, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	result := GenerateArt(input, banner)
	fmt.Print(result)
}
