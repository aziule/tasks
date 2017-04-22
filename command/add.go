package command

import (
	"fmt"
	"github.com/aziule/tasks/storage"
	"github.com/aziule/tasks/task"
	"strings"
)

type AddCommand struct {
}

func (c *AddCommand) GetName() string {
	return "add"
}

func (c *AddCommand) Execute(args []string) error {
	fmt.Println("Add a task")
	err, t := task.NewTask(strings.Join(args, " "))

	if err != nil {
		return err
	}

	storage.CreateTask(t)

	return nil
}

func init() {
	RegisterCommand(&AddCommand{})
}
