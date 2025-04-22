package main

import (
	"datapod-for-react-go-json/developer/learn"
	"datapod-for-react-go-json/qtools/qcli"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		qcli.Message("Usage: go run main.go <exerciseNumber>", "info")
		qcli.Message("Example: go run main.go 001", "info")
		return
	}

	exerciseNumber := os.Args[1]

	functions := map[string]func(){
		"001;Working with map[string][string]": learn.Ex001,
		"002;Design Pattern: Builder": learn.Ex002,
	}

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
		return
	}

	if title != "" {
		qcli.Message(fmt.Sprintf("EX%s: %s", exerciseNumber, title), "star")
	} else {
		qcli.Message(fmt.Sprintf("EX%s", exerciseNumber), "star")
	}
	fn()
}
