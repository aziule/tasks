package command

import (
	"errors"
	"strconv"
	"github.com/aziule/tasks/task"
	"github.com/aziule/tasks/storage"
)

type EditCommand struct {
}

func (c *EditCommand) GetName() string {
	return "edit"
}

func (c *EditCommand) Execute(args []string) error {
	task, err := parseArgs(args)

	if err != nil {
		return errors.New("Invalid arguments provided")
	}

	err = storage.Update(&task)

	if err != nil {
		return err
	}

	return nil
}

func parseArgs(args []string) (task.Task, error) {
	var task task.Task

	if len(args) != 2 {
		return task, errors.New("Invalid number of arguments")
	}

	taskId, err := strconv.Atoi(args[0])

	if err != nil {
		return task, err
	}

	task.Id = taskId
	task.Text = args[1]

	return task, nil
}

func init() {
	RegisterCommand(&EditCommand{})
}
