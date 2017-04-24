package task

import (
	"errors"
	"fmt"
)

type Task struct {
	Id int
	Text string
}

func NewTask(text string) (error, *Task) {
	if text == "" {
		return errors.New(fmt.Sprintf("The task cannot be empty")), nil
	}

	id := 42

	return nil, &Task{id, text}
}
