package command

import (
	"github.com/aziule/tasks/storage"
	"fmt"
	"github.com/aziule/tasks/task"
)

type TodoCommand struct {
}

func (c *TodoCommand) GetName() string {
	return "todo"
}

func (c *TodoCommand) Execute(args []string) error {
	tasks, err := storage.GetByStatus(task.STATUS_TODO)

	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Println(t.ToString())
	}

	return nil
}

func init() {
	RegisterCommand(&TodoCommand{})
}
