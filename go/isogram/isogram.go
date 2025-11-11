package isogram

import "unicode"

func IsIsogram(word string) bool {
	m := make(map[rune]int)
	for _, v := range word {
		if unicode.IsLetter(v) {
			m[unicode.ToLower(v)] += 1
		}
	}

	for _, v := range m {
		if v > 1 {
			return false
		}
	}

	return true
}
