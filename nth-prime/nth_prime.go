package prime

import "errors"

func generateInts(n int) []int {
	ints := []int{}
	for i := 0; i < n; i++ {
		ints = append(ints, i+1)
	}
	return ints
}

func getPrimes(n int) []int {
	primes := []int{}
	nints := generateInts(n)

	for i := 1; i < n; i++ {
		pnum := i + 1
		if (pnum * pnum) > n {
			break
		}
		p := nints[i]
		if p > 1 {
			for c := i; c < n; c++ {
				result := p * (c + 1)
				if result > n {
					break
				}
				nints[result-1] = 0
			}
		}
	}

	for _, num := range nints {
		if num > 1 {
			primes = append(primes, num)
		}
	}
	return primes

}

func Nth(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("there is no zeroth prime")
	}
	rangeSize := 2 * n
	for {
		primes := getPrimes(rangeSize)
		if len(primes) >= n {
			return primes[n-1], nil
		}
		rangeSize += n
	}

}
