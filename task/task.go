package task

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id int
	Text string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(id int, text string) (error, *Task) {
	if text == "" {
		return errors.New(fmt.Sprintf("The task cannot be empty")), nil
	}

	return nil, &Task{id, text, time.Now(), time.Now()}
}
