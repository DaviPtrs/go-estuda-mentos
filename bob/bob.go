// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func containsAlpha(input string) bool {
	for _, r := range input {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func checkAllCapital(input string) bool {
	return containsAlpha(input) && input == strings.ToUpper(input)
}

func isQuestion(input string) bool {
	return strings.HasSuffix(input, "?")
}

func isSilence(input string) bool {
	return input == ""
}

func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	is_question := isQuestion(trimmed)

	if isSilence(trimmed) {
		return "Fine. Be that way!"
	}

	if checkAllCapital(trimmed) {
		if is_question {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if is_question {
		return "Sure."
	}

	return "Whatever."
}
