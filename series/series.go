package series

func All(n int, s string) []string {
	output := []string{}
	for i := 0; i <= (len(s) - n); i++ {
		limit := i + n
		output = append(output, s[i:limit])
	}
	return output
}

func UnsafeFirst(n int, s string) string {
	if len(s) < n {
		return ""
	}
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
