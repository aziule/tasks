package command

import (
	"fmt"
	"errors"
)

var commands []Command

func RegisterCommand(c Command) {
	commands = append(commands, c)
}

func HandleCommand(name string) error {
	for _, c := range commands {
		if c.GetName() == name {
			return c.Execute()
		}
	}

	return errors.New(fmt.Sprintf("Command \"%s\" does not exist", name))
}