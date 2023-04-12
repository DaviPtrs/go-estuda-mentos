package anagram

import (
	"sort"
	"strings"
)

func sortString(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool {
		return s[i] < s[j]
	})
	return string(s)
}

func isAnagram(target string, candidate string) bool {
	target = strings.ToLower(target)
	candidate = strings.ToLower(candidate)

	if target == candidate {
		return false
	}
	return sortString(target) == sortString(candidate)
}

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}

	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
