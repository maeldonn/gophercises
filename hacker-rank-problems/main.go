package main

import "unicode"

func camelcase(s string) int32 {
	words := int32(1)
	for _, char := range s {
		if unicode.IsUpper(char) {
			words++
		}
	}
	return words
}

func caesarCipher(s string, k int32) string {
	var shifted string
	for _, char := range s {
		if unicode.IsLetter(char) {
			shift := 'a'
			if unicode.IsUpper(char) {
				shift = 'A'
			}

			s := (char + k - shift) % rune(26)
			shifted += string(s + shift)
		} else {
			shifted += string(char)
		}
	}
	return shifted
}
