package rotationalcipher

func getCipheredRune(char rune, shiftKey int) rune {
	base := 0
	if char >= 'A' && char <= 'Z' {
		base = int('A')
	}
	if char >= 'a' && char <= 'z' {
		base = int('a')
	}

	if base == 0 {
		return char
	}

	return rune(((int(char) - base + shiftKey) % 26) + base)

}

func RotationalCipher(plain string, shiftKey int) string {
	var output string

	for _, r := range plain {
		output += string(getCipheredRune(r, shiftKey))
	}
	return output
}
