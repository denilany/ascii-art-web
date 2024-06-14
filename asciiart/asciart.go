package asciiart

import (
	"fmt"
	"log"

	"asciiweb/extension"
	"asciiweb/printart"
	"asciiweb/read"
)

func Ascii() {
	fileExt := extension.Ext("standard.txt")
	if !fileExt {
		fmt.Println("wrong file extension")
		return
	}

	readFile, err := read.ReadAscii("thinkertoy.txt")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	art := printart.AsciiArt(readFile, "hello")

	fmt.Print(art)
}
