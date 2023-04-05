package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	s := strings.ReplaceAll(id, " ", "")
	if len(s) <= 1 {
		return false
	}

	sum := 0
	parity := len(s) % 2
	for i, digit := range s {
		if unicode.IsNumber(digit) {
			parsed_digit := int(digit - '0')
			if i%2 == parity {
				parsed_digit *= 2
				if parsed_digit > 9 {
					parsed_digit -= 9
				}
			}
			sum += parsed_digit
		} else {
			return false
		}
	}
	return sum%10 == 0
}
