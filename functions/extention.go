package functions

import (
	"path/filepath"
)

func Ext(banner string) bool {
	file := filepath.Ext(banner)

	return file == ".txt"
}
