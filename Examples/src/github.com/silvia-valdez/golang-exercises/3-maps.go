package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of the counts of each “word” in the string s.
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	parts := strings.Split(s, " ")

	for _, v := range parts {
		result[v] = result[v] + 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
