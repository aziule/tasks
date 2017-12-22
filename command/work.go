package command

import (
	"errors"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/aziule/tasks/storage"
)

type WorkCommand struct {
}

func (c *WorkCommand) GetName() string {
	return "work"
}

func (c *WorkCommand) Execute(args []string) error {
	task, err := GetTaskFromArgs(args)

	if err != nil {
		return errors.New("Invalid arguments provided")
	}

	terminationWaiter := sync.WaitGroup{}
	terminationWaiter.Add(1)

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Counting time worked on task", task.Id)
	start := time.Now()
	minutesSpent := 0

	go func() {
		<-ch
		end := time.Now()

		seconds := end.Unix() - start.Unix()
		minutesSpent = int(math.Ceil(float64(seconds) / 60))
		terminationWaiter.Done()
	}()

	terminationWaiter.Wait()
	task.MinutesSpent += minutesSpent

	storage.Update(task)

	fmt.Printf("%d minute(s) spent working, total is now %d minute(s)", minutesSpent, task.MinutesSpent)

	return nil
}

func init() {
	RegisterCommand(&WorkCommand{})
}
