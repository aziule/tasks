package command

import (
	"github.com/aziule/tasks/task"
	"errors"
	"strconv"
	"github.com/aziule/tasks/storage"
)

type Command interface {
	GetName() string
	Execute(args []string) error
}

func GetTaskFromArgs(args []string) (*task.Task, error) {
	if len(args) < 1 {
		return nil, errors.New("Invalid number of arguments")
	}

	taskId, err := strconv.Atoi(args[0])

	if err != nil {
		return nil, err
	}

	return storage.FindById(taskId)
}
