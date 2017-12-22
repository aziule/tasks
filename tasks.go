package main

import (
	"fmt"
	"os"

	"github.com/aziule/tasks/command"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command to execute")
		os.Exit(1)
	}

	err := command.HandleCommand(os.Args[1], os.Args[2:])

	if err != nil {
		fmt.Println(err)
	}
}
