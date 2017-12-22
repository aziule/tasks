package command

import (
	"errors"
	"fmt"

	"github.com/aziule/tasks/storage"
)

type DoCommand struct {
}

func (c *DoCommand) GetName() string {
	return "do"
}

func (c *DoCommand) Execute(args []string) error {
	task, err := GetTaskFromArgs(args)

	if err != nil {
		return errors.New("Invalid arguments provided")
	}

	task.Do()

	err = storage.Update(task)

	if err != nil {
		return err
	}

	fmt.Printf("Task %d done\n", task.Id)

	return nil
}

func init() {
	RegisterCommand(&DoCommand{})
}
