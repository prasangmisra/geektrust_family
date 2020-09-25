package main

import (
	"fmt"
	"geektrust/actions"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Need 1 argument for the file")
		return
	}
	if len(args) > 2 {
		fmt.Println("Need only 1 argument for the file")
		return
	}
	initBool := actions.Init(args[1])
	if !initBool {
		return
	}
}
