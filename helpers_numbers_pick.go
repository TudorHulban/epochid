package epochid

import "strings"

func pickNumbersFrom(text string, howMany uint) string {
	var digitsFound uint
	var result strings.Builder

	for i := len(text) - 1; i >= 0 && digitsFound < howMany; i-- {
		b := text[i]

		if b >= '0' && b <= '9' {
			result.WriteByte(b)

			digitsFound++
		}
	}

	for digitsFound < howMany {
		result.WriteByte('0')

		digitsFound++
	}

	return result.String()
}
