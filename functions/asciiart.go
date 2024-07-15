package functions

import (
	"fmt"
	"strings"
)

const (
	asciiHeight = 8
)

func asciiArt(bannerSlice []string, input string) (string, error) {
	var result strings.Builder

	input = replaceUnprint(input)
	input = replace(input)
	for _, ch := range input {
		if (ch < 32 || ch > 126) && string(ch) != "\r" && string(ch) != "\n" {
			return "", fmt.Errorf("contains a non-printable character")
		}
	}

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
	return result.String(), nil
}
