package cmdline

import (
	"errors"
	"os"

	"github.com/reiyuchan/workshopdl/lib/util"
)

func RunSubCommand(commands []Command, args []string) {
	if len(args) > 0 {
		command, err := FindCommandByName(commands, args[0])
		if err == nil {
			command.Run(args[1:])
			return
		}
	}
	util.ExitError(os.ErrInvalid, true)
}

func FindCommandByName(commands []Command, commandName string) (Command, error) {
	for _, c := range commands {
		for _, n := range c.Names {
			if n == commandName {
				return c, nil
			}
		}
	}
	return Command{}, errors.New("command not found")
}
