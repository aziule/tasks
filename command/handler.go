package command

import (
	"errors"
	"fmt"
)

var commands []Command

func RegisterCommand(c Command) {
	commands = append(commands, c)
}

func HandleCommand(name string, args []string) error {
	for _, c := range commands {
		if c.GetName() == name {
			return c.Execute(args)
		}
	}

	return errors.New(fmt.Sprintf("Command \"%s\" does not exist", name))
}
