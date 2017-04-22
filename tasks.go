package main

import (
	"os"
	"fmt"
	"github.com/aziule/tasks/command"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command to execute")
		os.Exit(1)
	}

	err := command.HandleCommand(os.Args[1])

	if err != nil {
		fmt.Println(err)
	}
}
