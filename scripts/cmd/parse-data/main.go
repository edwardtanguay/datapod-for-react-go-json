package main

import (
	"fmt"
)

func main() {

	devlog("this will parse the data")
	lines := getLinesFromFile("../../../data/flashcards.txt")

	for i, line := range lines {
		fmt.Printf("%d: WORKS: %s\n", i+1, line)
	}
}
