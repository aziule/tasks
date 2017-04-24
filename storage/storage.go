package storage

import (
	"fmt"
	"os"
	"encoding/csv"
	"github.com/aziule/tasks/task"
	"strconv"
	"io"
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

func Update(t *task.Task) error {
	fmt.Println("Should update task", t)
	return nil
}

func Add(t *task.Task) error {
	file, err := os.OpenFile(FILE_NAME, os.O_RDWR | os.O_APPEND, 0660)

	defer file.Close()

	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if t.Id == 0 {
		taskId, err := nextId()

		if err != nil {
			return err
		}

		t.Id = taskId
	}

	if err := writer.Write(toCsvRecord(t)); err != nil {
		return err
	}

	return nil
}

func nextId() (int, error) {
	file, err := os.Open(FILE_NAME)

	defer file.Close()

	if err != nil {
		return 0, err
	}

	reader := csv.NewReader(file)
	lastRecord := []string{}

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		lastRecord = record
	}

	if len(lastRecord) == 0 {
		return 1, nil
	}

	lastId, _ := strconv.Atoi(lastRecord[0])

	return lastId + 1, nil
}

func toCsvRecord(t *task.Task) []string {
	return []string{
		strconv.Itoa(t.Id),
		t.Text,
	}
}
