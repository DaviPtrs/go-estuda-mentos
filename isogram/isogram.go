package isogram

import "strings"

func IsIsogram(word string) bool {
	runeMap := make(map[rune]int)
	for _, r := range strings.ToLower(word) {
		_, test := runeMap[r]
		if !test || r == ' ' || r == '-' {
			runeMap[r] = 0
		} else {
			return false
		}
	}
	return true
}
