package main

import (
	"datapod-for-react-go-json/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Flashcard struct {
	Suuid    string `json:"suuid"`
	Category string `json:"category"`
	Front    string `json:"front"`
	Back     string `json:"back"`
}

func main() {
	utils.Devlog("parsing flashcards.txt into flashcards.json...")
	lines := utils.GetLinesFromFile("../../../data/flashcards.txt")

	var flashcards []Flashcard
	for i := 0; i < len(lines); i += 2 {
		if i+1 >= len(lines) {
			break
		}
		front := strings.TrimSpace(lines[i])
		back := strings.TrimSpace(lines[i+1])
		category := "general" // Default category
		if strings.Contains(front, ":") {
			parts := strings.SplitN(front, ":", 2)
			category = strings.TrimSpace(parts[0])
			front = strings.TrimSpace(parts[1])
		}

		flashcards = append(flashcards, Flashcard{
			Suuid:    utils.GenerateShortUUID(6),
			Category: category,
			Front:    front,
			Back:     back,
		})
	}

	jsonData, err := json.MarshalIndent(flashcards, "", "\t")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	err = os.WriteFile("../../../datajson/flashcards.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		return
	}

	fmt.Println("Successfully updated flashcards.json")
}
