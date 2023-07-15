package stringset

import "fmt"

type Set map[string]bool

func New() Set {
	set := map[string]bool{}
	return set
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, item := range l {
		set[item] = true
	}
	return set
}

func (s Set) String() string {
	output := "{"
	first := true
	for item := range s {
		if !first {
			output += ", "
		}
		first = false
		output += fmt.Sprintf("\"%s\"", item)
	}
	output += "}"
	return output
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, check := s[elem]
	return check
}

func (s Set) Add(elem string) {
	if !s.Has(elem) {
		s[elem] = true
	}
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for elem := range s1 {
		if s2.Has(elem) {
			return false
		}
	}
	for elem := range s2 {
		if s1.Has(elem) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		if s2.Has(elem) {
			result.Add(elem)
		}
	}

	return result
}

func Difference(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		if !s2.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	result := New()

	for elem := range s1 {
		result.Add(elem)
	}

	for elem := range s2 {
		result.Add(elem)
	}

	return result
}
