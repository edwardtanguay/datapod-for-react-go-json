package main

import (
	"datapod-for-react-go-json/cmd/utils"
	"fmt"
)

func main() {

	utils.Devlog("this will parse the data")
	lines := utils.GetLinesFromFile("../../../data/flashcards.txt")

	for i, line := range lines {
		fmt.Printf("%d: WORKS: %s\n", i+1, line)
	}
}
