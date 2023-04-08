// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

type Kind int

const (
	NaT = Kind(0)
	Equ = Kind(3)
	Iso = Kind(2)
	Sca = Kind(1)
)

func validateTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	sum1 := a + b
	sum2 := a + c
	sum3 := c + b
	if sum1 < c || sum2 < b || sum3 < a {
		return false
	}
	return true
}

func KindFromSides(a, b, c float64) Kind {
	if !validateTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
