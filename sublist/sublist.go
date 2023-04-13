package sublist

func checkListStartEqual(l1, l2 []int) bool {
	if len(l2) < len(l1) {
		return false
	}
	for i, elem1 := range l1 {
		if elem1 != l2[i] {
			return false
		}
	}
	return true
}

// is l1 sublist of l2?
func isSublist(l1, l2 []int) bool {
	if len(l1) == 0 {
		return true
	}
	for i, elem := range l2 {
		if elem == l1[0] {
			if (len(l2) - i) < len(l1) {
				return false
			}
			if checkListStartEqual(l1, l2[i:]) {
				return true
			}
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	lenl1 := len(l1)
	lenl2 := len(l2)

	// Equal check
	if lenl1 == lenl2 {
		if isSublist(l1, l2) {
			return RelationEqual
		}
	}

	// Superlist check
	if lenl1 > lenl2 {
		if isSublist(l2, l1) {
			return RelationSuperlist
		}
	}

	// Sublist check
	if lenl1 < lenl2 {
		if isSublist(l1, l2) {
			return RelationSublist
		}
	}

	return RelationUnequal

}
