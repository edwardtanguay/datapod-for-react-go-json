package main

import (
	"datapod-for-react-go-json/cmd/utils"
	"fmt"
)

func main() {

	utils.Devlog("printing flashcard file")
	lines := utils.GetLinesFromFile("../../../data/flashcards.txt")

	for i, line := range lines {
		fmt.Printf("%03d: %s\n", i+1, line)
	}
}
