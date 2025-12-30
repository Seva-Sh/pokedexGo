package main

import (
	"strings"
)

func cleanInput(text string) []string {
	loweredWords := strings.ToLower(text)
	splitWords := strings.Fields(loweredWords)
	return splitWords
}
