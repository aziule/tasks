package command

import (
	"errors"
	"strconv"
	"github.com/aziule/tasks/task"
	"github.com/aziule/tasks/storage"
	"fmt"
)

type EditCommand struct {
}

func (c *EditCommand) GetName() string {
	return "edit"
}

func (c *EditCommand) Execute(args []string) error {
	task, err := c.parseArgs(args)

	if err != nil {
		return errors.New("Invalid arguments provided")
	}

	err = storage.Update(task)

	if err != nil {
		return err
	}

	fmt.Printf("Task %d edited\n", task.Id)

	return nil
}

func (c *EditCommand) parseArgs(args []string) (*task.Task, error) {
	if len(args) != 2 {
		return nil, errors.New("Invalid number of arguments")
	}

	taskId, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, err
	}

	task, err := storage.FindById(taskId)

	if err != nil {
		return nil, err
	}

	task.Id = taskId
	task.Text = args[1]

	return task, nil
}

func init() {
	RegisterCommand(&EditCommand{})
}
