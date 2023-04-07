package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	filtered := Ints(nil)

	for _, elem := range i {
		if filter(elem) {
			filtered = append(filtered, elem)
		}
	}

	return filtered
}

func (i Ints) Discard(filter func(int) bool) Ints {
	filtered := Ints(nil)

	for _, elem := range i {
		if !filter(elem) {
			filtered = append(filtered, elem)
		}
	}

	return filtered
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	filtered := Lists(nil)

	for _, elem := range l {
		if filter(elem) {
			filtered = append(filtered, elem)
		}
	}

	return filtered

}

func (s Strings) Keep(filter func(string) bool) Strings {
	filtered := Strings(nil)

	for _, elem := range s {
		if filter(elem) {
			filtered = append(filtered, elem)
		}
	}

	return filtered

}
