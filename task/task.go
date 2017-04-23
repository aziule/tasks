package task

import (
	"errors"
	"fmt"
	"strconv"
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

func (t *Task) ToStringSlice() []string {
	return []string{
		strconv.Itoa(t.Id),
		t.Text,
	}
}