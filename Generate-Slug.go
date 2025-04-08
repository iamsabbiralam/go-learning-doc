package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	message := "Good Morning"
	slug := GenerateSlug(message)
	fmt.Println(slug)
}

func GenerateSlug(input string) string {
	if input == "" {
		return ""
	}

	// Convert to lowercase and trim whitespace
	slug := strings.ToLower(strings.TrimSpace(input))

	// Replace all non-alphanumeric characters with dashes
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")

	// Trim any leading or trailing dashes
	slug = strings.Trim(slug, "-")

	return slug
}
