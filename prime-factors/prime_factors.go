package prime

import (
	"math"
)

func Factors(n int64) []int64 {
	rest := n
	limit := int64(math.Sqrt(float64(n)))
	if limit < n {
		limit = n
	}
	primes := []int64{}
	for c := int64(2); c < limit; c++ {
		if rest <= 1 {
			break
		}
		if rest%c == 0 {
			rest /= c
			primes = append(primes, c)
			c--
		}
	}
	return primes
}
