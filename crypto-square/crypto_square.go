package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func findCValue(length int) int {
	start := int(math.Sqrt(float64(length)))
	num := start
	result := 0
	for {
		result = num * num
		if result < length {
			num++
		} else {
			break
		}
	}
	return num
}

func findRValue(length int, c int) int {
	if c*(c-1) >= length {
		return c - 1
	}
	return c
}

func filter(str string) string {
	new := []rune{}
	str = strings.ToLower(str)
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			new = append(new, r)
		}
	}
	return string(new)
}

func initializeMatrix(r int, c int) [][]rune {
	out := make([][]rune, r)
	for i := range out {
		out[i] = make([]rune, c)
		for j := 0; j < c; j++ {
			out[i][j] = rune(32) // Filling slots with white space
		}
	}
	return out
}

func transposeMatrix(original [][]rune, r int, c int) [][]rune {
	transposed := initializeMatrix(c, r)
	for i, row := range original {
		for j := range row {
			transposed[j][i] = original[i][j]
		}
	}
	return transposed
}

func encodeMatrix(mtx [][]rune, r int, c int) string {
	tranposed := transposeMatrix(mtx, r, c)
	output := ""
	for i, row := range tranposed {
		output += string(row)
		if i != c-1 {
			output += " "
		}
	}
	return output
}

func Encode(pt string) string {
	pt = filter(pt)
	length := len(pt)
	if length == 0 {
		return pt
	}
	c := findCValue(length)
	r := findRValue(length, c)
	auxMatrix := initializeMatrix(r, c)

	i := 0
	j := 0
	for _, char := range pt {
		auxMatrix[i][j] = char
		if j != (c - 1) {
			j++
		} else {
			i++
			j = 0
		}
	}

	return encodeMatrix(auxMatrix, r, c)

}
