package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aziule/tasks/storage"
	"github.com/aziule/tasks/task"
)

type AddCommand struct {
}

func (c *AddCommand) GetName() string {
	return "add"
}

func (c *AddCommand) Execute(args []string) error {
	err, t := task.NewTask(0, strings.Join(args, " "))

	if err != nil {
		return errors.New("An error occured when creating the task")
	}

	err = storage.Add(t)

	if err != nil {
		return errors.New("An error occured when saving the task")
	}

	fmt.Printf("Task %d created\n", t.Id)

	return nil
}

func init() {
	RegisterCommand(&AddCommand{})
}
