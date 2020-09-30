package main

import (
	"fmt"
	"geektrust/actions"
	"geektrust/helpers"
	"os"
)

func main() {
	args := os.Args
	switch true {
	case len(args) < 2:
		fmt.Println("Need 1 argument for the file")
		return
	case len(args) > 2:
		fmt.Println("Need only 1 argument for the file")
		return
	default:
		isFileExist := helpers.FileExists(args[1])
		if isFileExist {
			actions.Init(args[1])
		}

	}
	// initBool :=

	// if !initBool {
	// 	return
	// }
}
