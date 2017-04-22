package command

import (
	"github.com/aziule/tasks/storage"
	"github.com/aziule/tasks/task"
	"strings"
	"errors"
)

type AddCommand struct {
}

func (c *AddCommand) GetName() string {
	return "add"
}

func (c *AddCommand) Execute(args []string) error {
	err, t := task.NewTask(strings.Join(args, " "))

	if err != nil {
		return errors.New("An error occured when creating the task")
	}

	err = storage.CreateTask(t)

	if err != nil {
		return errors.New("An error occured when saving the task")
	}

	return nil
}

func init() {
	RegisterCommand(&AddCommand{})
}
