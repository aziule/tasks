package task

import (
	"errors"
	"fmt"
)

type Task struct {
	Id int
	Text string
}

func NewTask(id int, text string) (error, *Task) {
	if text == "" {
		return errors.New(fmt.Sprintf("The task cannot be empty")), nil
	}

	return nil, &Task{id, text}
}
