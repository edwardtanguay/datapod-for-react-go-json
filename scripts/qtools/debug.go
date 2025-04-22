package qtools

import "fmt"

/*
Output information in console in a uniform way

devlog("no files are locked")

devlog(fmt.Sprintf("There are %d flashcards.", len(flashcards)))
*/
func Debug(line string) {
	fmt.Printf("DEBUG ### %s ################################\n", line)
}
