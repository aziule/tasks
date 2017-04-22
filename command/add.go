package command

import "fmt"

type AddCommand struct {
}

func (c *AddCommand) GetName() string {
	return "add"
}

func (c *AddCommand) Execute() error {
	fmt.Println("Add a task")
	return nil
}

func init() {
	RegisterCommand(&AddCommand{})
}