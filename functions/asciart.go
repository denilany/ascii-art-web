package functions

import (
	"fmt"
	"log"
)

func Ascii() {
	fileExt := Ext("standard.txt")
	if !fileExt {
		fmt.Println("wrong file extension")
		return
	}

	readFile, err := ReadAscii("thinkertoy.txt")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	art := AsciiArt(readFile, "hello")

	fmt.Print(art)
}
