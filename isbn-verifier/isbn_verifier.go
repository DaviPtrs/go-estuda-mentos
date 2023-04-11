package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	length := len(isbn)
	if length != 10 {
		return false
	}

	formula := 0
	for i, digit := range isbn {
		if !unicode.IsDigit(digit) && (i != (length-1) || digit != 'X') {
			return false
		}
		digitValue := 0
		if digit == 'X' {
			digitValue = 10
		} else {
			digitValue = int(digit) - '0'
		}

		formula += digitValue * (length - i)
	}

	return formula%11 == 0
}
