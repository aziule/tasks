package storage

import (
	"fmt"
	"github.com/aziule/tasks/task"
	"os"
	"encoding/csv"
	"strconv"
)

const FILE_NAME = "tasks.csv"

func init() {
	if _, err := os.Stat(FILE_NAME); err != nil {
		file, err := os.Create(FILE_NAME);

		if err != nil {
			fmt.Println("Impossible to create the file", FILE_NAME)
			os.Exit(1)
		}

		defer file.Close()
	}
}

func CreateTask(t *task.Task) error {
	file, err := os.OpenFile(FILE_NAME, os.O_RDWR | os.O_APPEND, 0660)

	defer file.Close()

	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(taskToCsv(t)); err != nil {
		return err
	}


	return nil
}

func taskToCsv(t *task.Task) []string {
	return []string{
		strconv.Itoa(t.Id),
		t.Text,
	}
}