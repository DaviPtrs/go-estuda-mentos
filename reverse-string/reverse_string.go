package reverse

func Reverse(input string) string {
	rns := []rune(input)

	rlen := len(rns)
	for i := 0; i < (rlen / 2); i++ {
		aux := rns[i]
		changeIdx := rlen - 1 - i
		rns[i] = rns[changeIdx]
		rns[changeIdx] = aux
	}
	return string(rns)
}
