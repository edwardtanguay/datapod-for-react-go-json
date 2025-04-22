package qdev

import (
	"fmt"
	"time"
)

/*
Output information in console in a uniform way

devlog("no files are locked")

devlog(fmt.Sprintf("There are %d flashcards.", len(flashcards)))
*/

// Debug outputs information to the console in a uniform way.
func Debug(obj interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("🛠️  %s - %v\n", timestamp, obj)
}
