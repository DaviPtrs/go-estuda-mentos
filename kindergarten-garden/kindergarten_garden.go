package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	kids   []string
	plants [2][]rune
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getPlant(id rune) string {
	switch id {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	default:
		return ""
	}
}

func (g *Garden) validateGarden() bool {
	for _, row := range g.plants {
		for _, cup := range row {
			if getPlant(cup) == "" {
				return false
			}
		}
	}
	return true
}

func getIndex(elems []string, elem string) int {
	for i, x := range elems {
		if x == elem {
			return i
		}
	}
	return -1
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	err := errors.New("invalid diagram")
	if len(diagram) == 0 || !strings.HasPrefix(diagram, "\n") {
		return nil, err
	}

	rows := strings.Fields(diagram)
	children = removeDuplicateStr(children)
	if len(rows) != 2 || len(rows[0]) != (2*len(children)) || len(rows[0]) != len(rows[1]) {
		return nil, err
	}
	obj := new(Garden)
	obj.plants[0] = []rune(rows[0])
	obj.plants[1] = []rune(rows[1])
	if !obj.validateGarden() {
		return nil, err
	}
	sort.Strings(children)
	obj.kids = children

	return obj, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants := []string{}
	kidIdx := getIndex(g.kids, child)
	if kidIdx == -1 {
		return nil, false
	}
	start := kidIdx * 2

	for _, row := range g.plants {
		for i := start; i < (start + 2); i++ {
			plants = append(plants, getPlant(row[i]))
		}
	}
	return plants, true
}
