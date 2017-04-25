package command

import (
	"github.com/aziule/tasks/storage"
	"fmt"
)

type ListCommand struct {
}

func (c *ListCommand) GetName() string {
	return "list"
}

func (c *ListCommand) Execute(args []string) error {
	tasks, err := storage.GetAll()

	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Println(t.Id, t.Text, t.UpdatedAt.Format("02/01/2006 3:04PM"))
	}

	return nil
}

func init() {
	RegisterCommand(&ListCommand{})
}
