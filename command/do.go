package command

import (
	"errors"
	"strconv"
	"github.com/aziule/tasks/task"
	"github.com/aziule/tasks/storage"
	"fmt"
)

type DoCommand struct {
}

func (c *DoCommand) GetName() string {
	return "do"
}

func (c *DoCommand) Execute(args []string) error {
	task, err := c.parseArgs(args)

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

func (c *DoCommand) parseArgs(args []string) (*task.Task, error) {
	if len(args) != 1 {
		return nil, errors.New("Invalid number of arguments")
	}

	taskId, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, err
	}

	return storage.FindById(taskId)
}

func init() {
	RegisterCommand(&DoCommand{})
}
