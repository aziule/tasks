package command

import (
	"errors"
	"github.com/aziule/tasks/storage"
	"fmt"
	"strings"
)

type EditCommand struct {
}

func (c *EditCommand) GetName() string {
	return "edit"
}

func (c *EditCommand) Execute(args []string) error {
	if len(args) < 2 {
		return errors.New("Invalid number of arguments")
	}

	task, err := GetTaskFromArgs(args)

	if err != nil {
		return errors.New("Invalid arguments provided")
	}

	task.Text = strings.Join(args[1:], " ")

	err = storage.Update(task)

	if err != nil {
		return err
	}

	fmt.Printf("Task %d edited\n", task.Id)

	return nil
}

func init() {
	RegisterCommand(&EditCommand{})
}
