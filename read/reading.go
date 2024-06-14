package read

import (
	"fmt"
	"os"
	"strings"
)

func ReadAscii(banner string) ([]string, error) {
	file, err := os.ReadFile(banner)
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	var splitData []string
	if banner == "thinkertoy.txt" {
		splitData = strings.Split(string(file), "\r\n")
	} else {
		splitData = strings.Split(string(file), "\n")
	}
	return splitData, nil
}
