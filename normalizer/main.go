package main

import (
	"strings"
	"unicode"
)

func normalize(phone string) string {
	var builder strings.Builder
	for _, ch := range phone {
		if unicode.IsNumber(ch) {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
