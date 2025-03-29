package main

import (
	"fmt"
	"os"
	"strings"
)

func devlog(line string) {
	fmt.Printf("2>>> %s\n", line)
}

/*
Get all lines from a file as a slice of strings

lines := getLinesFromFile("../../notes.txt")

- use relative path
*/
func getLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}
