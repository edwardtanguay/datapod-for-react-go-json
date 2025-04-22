package main

import (
	"datapod-for-react-go-json/developer/learn"
	"datapod-for-react-go-json/qtools/qcli"
	"fmt"
	"os"
	"strings"
)

func displayExercises(exercises map[string]func()) {
	var builder strings.Builder
	builder.WriteString("\nAvailable exercises:\n")

	for key := range exercises {
		parts := strings.SplitN(key, ";", 2)
		if len(parts) > 1 {
			builder.WriteString(fmt.Sprintf("  %s: %s\n", parts[0], parts[1]))
		} else {
			builder.WriteString(fmt.Sprintf("  %s\n", parts[0]))
		}
	}

	qcli.Message(builder.String(), "info")
	qcli.Message("Usage: go run main.go <exerciseNumber>", "info")
	qcli.Message("Example: go run main.go 001", "info")
}

func main() {

	functions := map[string]func(){
		"001;Working with map[string][string]": learn.Ex001,
		"002;Design Pattern: Builder":          learn.Ex002,
	}

	if len(os.Args) < 2 {
		displayExercises(functions)
		return
	}

	exerciseNumber := os.Args[1]

	var fn func()
	var title string
	exists := false

	for key, f := range functions {
		parts := strings.SplitN(key, ";", 2)
		if parts[0] == exerciseNumber {
			fn = f
			exists = true
			if len(parts) > 1 {
				title = parts[1]
			}
			break
		}
	}

	if !exists {
		qcli.Message(fmt.Sprintf("Exercise number is not valid: %s\n", exerciseNumber), "error")
		displayExercises(functions)
		return
	}

	if title != "" {
		msg := fmt.Sprintf("EX%s: %s", exerciseNumber, title)
		qcli.Message(msg, "star")
		fmt.Println(strings.Repeat("=", len(msg) + 3)) 
	} else {
		qcli.Message(fmt.Sprintf("EX%s", exerciseNumber), "star")
	}
	fn()
}
