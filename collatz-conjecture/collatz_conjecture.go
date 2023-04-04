package collatzconjecture

import "errors"

func isEven(n int) bool {
	return n%2 == 0
}

func CollatzConjecture(n int) (int, error) {
	i := 0
	if n <= 0 {
		return 0, errors.New("zero or negative is not permited")
	}
	for _ = 0; n != 1; i++ {
		if isEven(n) {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}
	return i, nil
}
