package sequence

import (
	"strings"
)

func ReplaceUnprint(input string) string {
	char := []string{"\\a", "\\r", "\\f", "\\v", "\\`", "\\x20"}

	for _, wrd := range char {
		if strings.Contains(input, wrd) {
			input = strings.ReplaceAll(input, wrd, "")
		}
	}

	return input
}
