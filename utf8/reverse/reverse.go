//go:build !solution

package reverse

func Reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i++ {
		runes[i], runes[j] = runes[j], runes[i]
		j--
	}

	return string(runes)
}
