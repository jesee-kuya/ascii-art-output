package ascii

import (
	"fmt"
)

// Methods sorts out the arguments recieved based on the number of arguments written and if the color flag has been passed
func (receive *Receiver) SortArg(argsPassed []string) string {
	var filename string
	msg := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
	if len(argsPassed) == 3 {
		if !IsFlagPassed("color") {
			return msg
		}
		filename = argsPassed[2]
		receive.ArgsPassed = argsPassed[:2]
	} else if len(argsPassed) == 2 {
		if IsFlagPassed("color") {
			_, err := GetFileName(argsPassed[1])
			if err != nil {
				filename = "standard"
				receive.ArgsPassed = argsPassed[:2]
			} else {
				filename = argsPassed[1]
				receive.ArgsPassed = argsPassed[:1]
			}
		} else {
			filename = argsPassed[1]
			receive.ArgsPassed = argsPassed[:1]
		}
	} else if len(argsPassed) == 1 {
		filename = "standard"
		receive.ArgsPassed = argsPassed[:1]
	}

	bannerContent, err := GetFileName(filename)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	receive.FileArr = bannerContent
	return ""
}
