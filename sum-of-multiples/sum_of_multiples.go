package summultiples

func checkMultiple(num int, divisors ...int) bool {
	for _, div := range divisors {
		if div != 0 {
			if num%div == 0 {
				return true
			}
		}
	}
	return false
}

func findFirstNonZero(numbers ...int) int {
	for _, num := range numbers {
		if num != 0 {
			return num
		}
	}
	return 1
}

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	start := findFirstNonZero(divisors...) - 1
	for i := start; i < limit-1; i++ {
		if checkMultiple(i+1, divisors...) {
			sum += i + 1
		}
	}

	return sum

}
