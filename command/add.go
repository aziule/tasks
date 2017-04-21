package command

import "fmt"

type AddCommand struct {
}

func (c *AddCommand) Execute() error {
	fmt.Println("Add a task")
	return nil
}

func (c *AddCommand) LoadFlags(args []string) error {
	return nil
}
