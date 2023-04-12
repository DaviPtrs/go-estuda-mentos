package cipher

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

type shift struct {
	distance int
}

type vigenere struct {
	key       string
	distances []int
}

func getCipheredRune(char rune, shiftKey int) rune {
	char = unicode.ToLower(char)
	if !unicode.IsLetter(char) {
		return char
	}
	result := (int(char) - 'a' + shiftKey) % 26
	if result < 0 {
		return rune('z' + result + 1)
	}
	return rune('a' + result)
}

func clearString(s string) string {
	newString := strings.ToLower(s)
	cleanRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	newString = cleanRegex.ReplaceAllString(newString, "")
	newString = strings.ReplaceAll(newString, " ", "")

	return newString
}

func shiftString(input string, distance int) string {
	ciphered := []rune{}
	input = clearString(input)
	for _, char := range input {
		if unicode.IsLetter(char) {
			ciphered = append(ciphered, getCipheredRune(char, distance))
		}
	}
	return string(ciphered)
}

func shiftStringFromArr(input string, distances []int, factor int) string {
	ciphered := []rune{}
	input = clearString(input)
	for i, char := range input {
		distance := (distances[i%len(distances)]) * factor
		if unicode.IsLetter(char) {
			ciphered = append(ciphered, getCipheredRune(char, distance))
		}
	}
	return string(ciphered)
}

func NewCaesar() Cipher {
	return shift{3}
}

func NewShift(distance int) Cipher {
	if math.Abs(float64(distance)) > 0 && math.Abs(float64(distance)) < 26 {
		return shift{distance}
	}
	return nil
}

func (c shift) Encode(input string) string {
	return shiftString(input, c.distance)
}

func (c shift) Decode(input string) string {
	return shiftString(input, -c.distance)
}

func NewVigenere(key string) Cipher {
	distances := []int{}
	invalid := true
	for _, char := range key {
		if !unicode.IsLetter(char) || !unicode.IsLower(char) {
			return nil
		}
		if invalid && char != 'a' {
			invalid = false
		}
		distances = append(distances, int(char-'a'))
	}
	if invalid {
		return nil
	}
	return vigenere{key, distances}
}

func (v vigenere) Encode(input string) string {
	return shiftStringFromArr(input, v.distances, 1)
}

func (v vigenere) Decode(input string) string {
	return shiftStringFromArr(input, v.distances, -1)
}
