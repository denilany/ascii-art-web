package printart

import (
	"fmt"
	"strings"

	"asciiweb/sequence"
)

const (
	asciiHeight = 8
)

func AsciiArt(bannerSlice []string, input string) string {
	var result strings.Builder

	input = sequence.ReplaceUnprint(input)
	input = sequence.Replace(input)

	arguments := strings.Split(input, "\r\n")

	for _, word := range arguments {
		if word == "" {
			fmt.Println()
		} else {
			for j := 0; j < asciiHeight; j++ {
				for _, ch := range word {
					index := int(ch-32)*9 + 1
					result.WriteString(bannerSlice[index+j])
				}
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
