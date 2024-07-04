package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	// Check for the correct format for the color flag
	for _, v := range os.Args {
		if v == "--color" || v == "-color" || v == "-output" || v == "--output" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		} else if strings.HasPrefix(v, "-color") || strings.HasPrefix(v, "-output") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	}
	var receive ascii.Receiver

	flag.StringVar(&receive.Colorflag, "color", "reset", "color for color input")
	flag.StringVar(&receive.Outputflag, "output", "banner.txt", "Write ascii output to a specified file")
	flag.Parse()
	argsPassed := flag.Args()

	// Methods sorts-out our arguments to the receive struct
	msg := receive.SortArg(argsPassed)
	if msg != "" {
		fmt.Println(msg)
		return
	}
	// Checks if the color flag has been passed
	if ascii.IsFlagPassed("color") {
		receive.Color()
	} else {
		receive.ColorCode = ""
		receive.Art()
	}
}
