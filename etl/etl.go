package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	newMap := make(map[string]int)
	for key, values := range in {
		for _, letter := range values {
			newMap[strings.ToLower(letter)] = key
		}
	}
	return newMap
}
