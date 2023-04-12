package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

func transpose(char rune) rune {
	if unicode.IsLetter(char) {
		return 'z' - char + 'a'
	} else {
		return char
	}
}

func clearString(s string) string {
	newString := strings.ToLower(s)
	cleanRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	newString = cleanRegex.ReplaceAllString(newString, "")
	newString = strings.ReplaceAll(newString, " ", "")

	return newString
}

func Atbash(s string) string {
	ciphered := ""
	s = clearString(s)

	for i, char := range s {
		transposed := transpose(char)
		ciphered += string(transposed)
		if (i+1)%5 == 0 && i != (len(s)-1) {
			ciphered += " "
		}
	}
	return ciphered
}
