package resistorcolorduo

import "strings"

func getColorValue(color string) int {
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
	default:
		return 0
	}
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) < 2 {
		return 0
	}

	value := 0
	for i := 0; i < 2; i++ {
		if i != 0 {
			value *= 10
		}
		value += getColorValue(strings.ToLower(colors[i]))
	}
	return value
}
