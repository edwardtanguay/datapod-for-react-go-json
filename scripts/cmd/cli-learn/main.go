package main

import (
	"datapod-for-react-go-json/developer/learn"
	"fmt"
	"os"
	// import "datapod-for-react-go-json/qtools/qcli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <exerciseNumber>")
		fmt.Println("Example: go run main.go 001")
		return
	}

	exerciseNumber := os.Args[1]

	switch exerciseNumber {
	case "001":
		learn.Ex001()
	case "002":
		learn.Ex002()
	default:
		fmt.Printf("Invalid exercise number: %s\n", exerciseNumber)
	}
}
