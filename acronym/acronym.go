// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "unicode"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	acr := []rune{}

	flag := true
	for _, char := range s {
		if flag && unicode.IsLetter(char) {
			acr = append(acr, unicode.ToUpper(char))
			flag = false
		}
		if unicode.IsSpace(char) || char == '-' {
			flag = true
		}
	}

	return string(acr)
}
