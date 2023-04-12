package prime

func Factors(n int64) []int64 {
	rest := n
	primes := []int64{}
	c := int64(2)
	for rest > 1 {
		if rest%c == 0 {
			rest /= c
			primes = append(primes, c)
		} else {
			c++
		}
	}
	return primes
}
