package sieve

func Sieve(limit int) []int {
	primes := []int{}
	nints := make([]bool, limit)
	nints[0] = true

	for i := 1; i < limit; i++ {
		p := i + 1
		if (p * p) > limit {
			break
		}
		check := nints[i]
		if !check {
			for c := i; c < limit; c++ {
				result := p * (c + 1)
				if result > limit {
					break
				}
				nints[result-1] = true
			}
		}
	}

	for i, num := range nints {
		if !num {
			primes = append(primes, i+1)
		}
	}
	return primes
}
