package main

import "strings"

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	strSlice := strings.Fields(loweredText)
	return strSlice
}
