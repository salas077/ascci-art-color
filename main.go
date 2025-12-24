package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if we have the correct number of arguments
	// We need at least 1 arg (the text to convert)
	// Maximum 2 args (text + optional banner name)
	if len(os.Args) < 2 || len(os.Args) > 3 {
		// Show helpful usage message instead of silent exit
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  go run . \"Hello\"")
		fmt.Println("  go run . \"Hello\" shadow")
		fmt.Println()
		fmt.Println("Available banners: standard, shadow, thinkertoy")
		return
	}

	// First argument is always the text we want to convert to ASCII art
	input := os.Args[1]

	// Default to standard banner if no banner is specified
	bannerName := "standard"
	if len(os.Args) == 3 {
		// Second argument is the banner choice (standard, shadow, or thinkertoy)
		bannerName = os.Args[2]
	}

	// Determine which banner file to load based on user's choice
	var bannerPath string
	switch bannerName {
	case "standard":
		bannerPath = "banners/standard.txt"
	case "shadow":
		bannerPath = "banners/shadow.txt"
	case "thinkertoy":
		bannerPath = "banners/thinkertoy.txt"
	default:
		// User provided an invalid banner name
		// Show error message with what they entered and what's available
		fmt.Printf("Error: Invalid banner '%s'\n", bannerName)
		fmt.Println("Available banners: standard, shadow, thinkertoy")
		return
	}

	// Try to load the selected banner file
	banner, err := LoadBanner(bannerPath)
	if err != nil {
		// Loading failed - show which file couldn't be loaded and why
		fmt.Printf("Error: Could not load banner file '%s'\n", bannerPath)
		fmt.Printf("Details: %v\n", err)
		return
	}

	// Convert the input text to ASCII art using the loaded banner
	output := RenderInput(input, banner)

	// Print the final ASCII art to the console
	fmt.Print(output)
}
