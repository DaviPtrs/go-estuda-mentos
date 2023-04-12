package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func checkKey(m Frequency, key string) bool {
	_, err := m[key]
	return err
}

func checkR(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func Split(phrase string) []string {
	rstring := []rune(phrase)
	for i := range rstring {
		r := rstring[i]
		if !checkR(r) {
			if r != '\'' || (i == 0 || i == len(rstring)-1 || !checkR(rstring[i-1]) || !checkR(rstring[i+1])) {
				rstring[i] = ' '
			}
		}
	}
	return strings.Fields(string(rstring))
}

func WordCount(phrase string) Frequency {
	freqmap := Frequency{}
	phrase = strings.ToLower(phrase)
	sslice := Split(phrase)

	for _, word := range sslice {
		if !checkKey(freqmap, word) {
			freqmap[word] = 0
		}
		freqmap[word]++
	}

	return freqmap
}
