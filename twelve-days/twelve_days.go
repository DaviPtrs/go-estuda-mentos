package twelve

import "fmt"

var objMap = map[int]string{
	1:  "Partridge in a Pear Tree",
	2:  "Turtle Doves",
	3:  "French Hens",
	4:  "Calling Birds",
	5:  "Gold Rings",
	6:  "Geese-a-Laying",
	7:  "Swans-a-Swimming",
	8:  "Maids-a-Milking",
	9:  "Ladies Dancing",
	10: "Lords-a-Leaping",
	11: "Pipers Piping",
	12: "Drummers Drumming",
}

func getExtenseNum(num int) (string, string) {
	switch num {
	case 1:
		return "first", ""
	case 2:
		return "second", "two"
	case 3:
		return "third", "three"
	case 4:
		return "fourth", "four"
	case 5:
		return "fifth", "five"
	case 6:
		return "sixth", "six"
	case 7:
		return "seventh", "seven"
	case 8:
		return "eighth", "eight"
	case 9:
		return "ninth", "nine"
	case 10:
		return "tenth", "ten"
	case 11:
		return "eleventh", "eleven"
	case 12:
		return "twelfth", "twelve"
	default:
		return "", ""
	}
}

func Verse(i int) string {
	if !(i >= 1 && i <= 12) {
		panic("Invalid number")
	}
	adj, _ := getExtenseNum(i)
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", adj)
	for c := i; c > 0; c-- {
		if c != 1 {
			_, name := getExtenseNum(c)
			verse += fmt.Sprintf(" %s %s,", name, objMap[c])
		} else {
			if i != 1 {
				verse += " and"
			}
			verse += fmt.Sprintf(" a %s.", objMap[1])
		}
	}

	return verse
}

func Song() string {
	song := ""
	for c := 1; c <= 12; c++ {
		song += Verse(c)
		if c != 12 {
			song += "\n"
		}
	}
	return song
}
