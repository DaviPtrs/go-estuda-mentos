package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	digits := []rune{}
	err := errors.New("invalid input")
	for _, char := range phoneNumber {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	length := len(digits)
	if length == 11 {
		if digits[0] != '1' {
			return "", err
		}
		digits = digits[1:]
		length--
	}

	if length != 10 {
		return "", err
	}

	for i, digit := range digits {
		if !(digit >= '0' && digit <= '9') {
			return "", err
		}
		if i == 0 || i == 3 {
			if digit < '2' {
				return "", err
			}
		}
	}

	return string(digits), nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err == nil {
		number = number[:3]
	}
	return number, err
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return number, err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
