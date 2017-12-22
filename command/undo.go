package command

import (
	"errors"
	"fmt"

	"github.com/aziule/tasks/storage"
)

type UndoCommand struct {
}

func (c *UndoCommand) GetName() string {
	return "undo"
}

func (c *UndoCommand) Execute(args []string) error {
	task, err := GetTaskFromArgs(args)

	if err != nil {
		return errors.New("Invalid arguments provided")
	}

	task.Undo()

	err = storage.Update(task)

	if err != nil {
		return err
	}

	fmt.Printf("Task %d set as being to do\n", task.Id)

	return nil
}

func init() {
	RegisterCommand(&UndoCommand{})
}
