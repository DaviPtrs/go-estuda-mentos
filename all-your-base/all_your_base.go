package allyourbase

import (
	"errors"
	"math"
)

func ReverseSlice(x []int) {
	for i := 0; i < len(x)/2; i++ {
		aux := x[i]
		x[i] = x[len(x)-i-1]
		x[len(x)-i-1] = aux
	}
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	number10 := 0
	for i, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		power := len(inputDigits) - i - 1
		number10 += digit * int(math.Pow(float64(inputBase), float64(power)))
	}
	if number10 == 0 {
		return []int{0}, nil
	}
	outputDigits := []int{}
	for number10 != 0 {
		outputDigits = append(outputDigits, int(number10%outputBase))
		number10 /= outputBase
	}
	ReverseSlice(outputDigits)
	return outputDigits, nil
}
