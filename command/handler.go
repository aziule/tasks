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
	eligibleCommands := []Command{}

	for _, c := range commands {
		if len(c.GetName()) >= len(name) && c.GetName()[:len(name)] == name {
			eligibleCommands = append(eligibleCommands, c)
		}
	}

	if len(eligibleCommands) == 1 {
		return eligibleCommands[0].Execute(args)
	}

	if len(eligibleCommands) > 1 {
		availableCommandsNames := []string{}

		for _, c := range eligibleCommands {
			availableCommandsNames = append(availableCommandsNames, c.GetName())
		}

		return errors.New(fmt.Sprintf("Multiple commands found: %s", availableCommandsNames))
	}

	return errors.New(fmt.Sprintf("Command \"%s\" does not exist", name))
}
