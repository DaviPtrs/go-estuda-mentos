package armstrong

import (
	"fmt"
	"math"
)

func toDigitSlice(n int) []int {
	output := []int{}
	str := fmt.Sprintf("%v", n)
	for _, digit := range str {
		output = append(output, int(digit-'0'))
	}
	return output
}

func IsNumber(n int) bool {
	digits := toDigitSlice(n)
	length := len(digits)

	sum := 0
	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(length)))
	}
	return sum == n
}
