package main

import (
	"datapod-for-react-go-json/developer/learn"
	"datapod-for-react-go-json/qtools/qcli"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		qcli.Message("Usage: go run main.go <exerciseNumber>", "info")
		qcli.Message("Example: go run main.go 001", "info")
		return
	}

	exerciseNumber := os.Args[1]

	functions := map[string]func(){
		"001": learn.Ex001,
		"002": learn.Ex002,
	}

	fn, exists := functions[exerciseNumber]
	if !exists {
		qcli.Message(fmt.Sprintf("Invalid exercise number: %s\n", exerciseNumber), "error")
		return
	}

	qcli.Message(fmt.Sprintf("Exercise Ex%s", exerciseNumber), "success")
	fn()
}
