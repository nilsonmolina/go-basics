// Package word provides functions for working with strings and words.
package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount returns the number of times each word is used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string.
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
