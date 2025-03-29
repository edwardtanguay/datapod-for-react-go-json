package main

import (
	"datapodforreactgojson/cmd/tools"
	"fmt"
)

func main() {

	tools.Devlog("this will parse the data")
	lines := tools.GetLinesFromFile("../../../data/flashcards.txt")

	for i, line := range lines {
		fmt.Printf("%d: WORKS: %s\n", i+1, line)
	}
}
