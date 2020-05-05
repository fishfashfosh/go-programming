// Package word provides utility functions for words
package word

import "strings"

// Return a map showing the words in the given string along with their counts
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count the number of words in the given string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
