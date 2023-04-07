package darts

import "math"

const outerCRadius int = 10
const middleCRadius int = 5
const innerCRadius int = 1

func Score(x, y float64) int {
	distance := math.Sqrt((x * x) + (y * y))
	if distance <= float64(innerCRadius) {
		return 10
	} else if distance <= float64(middleCRadius) {
		return 5
	} else if distance <= float64(outerCRadius) {
		return 1
	} else {
		return 0
	}
}
