package task

import (
	"errors"
	"fmt"
	"time"
)

const STATUS_TODO = "todo"
const STATUS_DONE = "done"

type Task struct {
	Id int
	Text string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(id int, text string) (error, *Task) {
	if text == "" {
		return errors.New(fmt.Sprintf("The task cannot be empty")), nil
	}

	return nil, &Task{id, text, STATUS_TODO, time.Now(), time.Now()}
}

func (t *Task) ToString() string {
	return fmt.Sprintf("%d %s %s", t.Id, t.Text, t.UpdatedAt.Format("02/01/2006 3:04PM"))
}

func (t *Task) Do() {
	t.Status = STATUS_DONE
}
