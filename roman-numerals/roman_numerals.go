package romannumerals

import (
	"errors"
)

func getRomanSymbol(num int) string {
	switch num {
	case 1000:
		return "M"
	case 900:
		return "CM"
	case 500:
		return "D"
	case 400:
		return "CD"
	case 100:
		return "C"
	case 90:
		return "XC"
	case 50:
		return "L"
	case 40:
		return "XL"
	case 10:
		return "X"
	case 9:
		return "IX"
	case 5:
		return "V"
	case 4:
		return "IV"
	case 1:
		return "I"
	default:
		return ""
	}
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid number")
	}

	possibleRomanNumbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	output := ""
	for _, number := range possibleRomanNumbers {
		for input >= number {
			output += getRomanSymbol(number)
			input -= number
		}
	}

	return output, nil

}
