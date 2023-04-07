package pangram

import "strings"

func IsPangram(input string) bool {
	lowerinput := strings.ToLower(input)

	for char := 'a'; char <= 'z'; char++ {
		if !strings.ContainsRune(lowerinput, rune(char)) {
			return false
		}
	}
	return true
}
