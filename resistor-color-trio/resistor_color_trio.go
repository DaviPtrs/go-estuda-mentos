package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").

func getColor(color string) int {
	switch color {
	case "black":
		return 0
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	}
	return -1
}

func Label(colors []string) string {
	var number int
	n1 := getColor(colors[0])
	if n1 != 0 {
		number = n1 * 10
	}
	number += getColor(colors[1])
	multiplier := getColor(colors[2])
	if number%10 == 0 && number != 0 {
		multiplier++
		number /= 10
	}
	sufix := ""
	if multiplier == 9 {
		multiplier -= 9
		sufix = "giga"
	}
	if multiplier >= 6 {
		multiplier -= 6
		sufix = "mega"
	}
	if multiplier >= 3 {
		multiplier -= 3
		sufix = "kilo"
	}
	number *= int(math.Pow10(multiplier))
	return fmt.Sprintf("%d %sohms", number, sufix)

}
