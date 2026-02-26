//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
)

func CollapseSpaces(input string) string {
	runes := []rune(input)
	skip := make([]bool, len(runes))

	for i, r := range runes {
		if unicode.IsSpace(r) && i+1 < len(runes) && unicode.IsSpace(runes[i+1]) {
			skip[i] = true
			continue
		}
	}

	// builder for efficiency
	var sb strings.Builder
	for i, r := range runes {
		if skip[i] {
			continue
		}

		if unicode.IsSpace(r) {
			sb.WriteRune(' ')
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
