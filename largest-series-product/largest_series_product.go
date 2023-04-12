package lsproduct

import (
	"errors"
	"unicode"
)

func runeToNum(r rune) int {
	return int(r) - '0'
}

func calcSpan(rdigits []rune, index int, span int) (int64, error) {
	var product int64 = 1
	for c := 0; c < span; c++ {
		rspan := rdigits[index+c]
		if !unicode.IsDigit(rspan) {
			return 0, errors.New("digits input must only contain digits")
		}
		product *= int64(runeToNum(rspan))
	}
	return product, nil
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}

	rdigits := []rune(digits)
	limit := len(rdigits) - span

	var product int64
	for i := 0; i <= limit; i++ {
		aux, err := calcSpan(rdigits, i, span)
		if err != nil {
			return aux, err
		}
		if aux > product {
			product = aux
		}
	}

	return product, nil
}
