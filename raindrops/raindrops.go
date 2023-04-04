package raindrops

import (
	"strconv"
)

func testFactor(number int, factor int) bool {
	return number%factor == 0
}

func Convert(number int) string {
	drop := ""
	if testFactor(number, 3) {
		drop += "Pling"
	}
	if testFactor(number, 5) {
		drop += "Plang"
	}
	if testFactor(number, 7) {
		drop += "Plong"
	}
	if drop == "" {
		drop = strconv.Itoa(number)
	}

	return drop
	// panic("Please implement the Convert function")
}
