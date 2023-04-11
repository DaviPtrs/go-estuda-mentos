// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var output []string
	length := len(rhyme)

	for i := range rhyme {
		var piece string
		if i == length-1 {
			piece = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			piece = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
		output = append(output, piece)
	}

	return output
}
