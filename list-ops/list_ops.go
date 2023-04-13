package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, item := range s {
		acc = fn(acc, item)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}
	for _, item := range s {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := s[:]

	for i, item := range s {
		result[i] = fn(item)
	}

	return result
}

func (s IntList) Reverse() IntList {
	result := s[:]
	for i := 0; i < int(len(s)/2); i++ {
		changeIdx := len(s) - 1 - i
		aux := result[i]
		result[i] = result[changeIdx]
		result[changeIdx] = aux
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := s[:]
	for _, item := range lst {
		result = append(result, item)
	}
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s[:]
	for _, lst := range lists {
		result = result.Append(lst)
	}
	return result
}
