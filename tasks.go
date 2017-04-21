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

	var c command.Command

	switch os.Args[1] {
		case "add":
			c = new(command.AddCommand)
		default:
			fmt.Println("Invalid command")
			os.Exit(1)
	}

	err := c.LoadFlags(os.Args[2:])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = c.Execute()

	if err != nil {
		panic(err)
	}
}
