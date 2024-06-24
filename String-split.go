package main

import (
	"fmt"
	"strings"
)

func extractUsername(email string) string {
	// Split the email at the '@' symbol
	parts := strings.Split(email, "@")
	// Return the first part which is the username
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func main() {
	email := "sabbiralam833@gmail.com"
	username := extractUsername(email)
	fmt.Println("Username:", username)
}
