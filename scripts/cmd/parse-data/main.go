package main

import (
	"datapod-for-react-go-json/utils"
	"fmt"
)

func main() {
	utils.Devlog("this shows the lines of the text file")
	lines := utils.GetLinesFromFile("../../../data/flashcards.txt")
	for i, line := range lines {
		fmt.Printf("%03d: %s\n", i+1, line)
	}
}
