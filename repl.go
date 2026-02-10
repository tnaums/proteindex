package main

import "strings"

// Function cleanInput returns a slice of strings
// where each is a 'field' or 'word' from the
// input string. Used to clean and parse repl input.
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
