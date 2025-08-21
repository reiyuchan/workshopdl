package main

import (
	"errors"
	"os"
	"runtime"

	"github.com/reiyuchan/workshopdl/get"
	"github.com/reiyuchan/workshopdl/help"
	"github.com/reiyuchan/workshopdl/lib/cmdline"
	"github.com/reiyuchan/workshopdl/lib/util"
)

func main() {
	if runtime.GOOS == "darwin" {
		util.ExitError(errors.ErrUnsupported, true)
	}

	commands := []cmdline.Command{
		get.GetCommand(),
	}

	helpCommand := help.GetCommand()
	commandsWithHelp := append(commands, cmdline.Command{
		Names: helpCommand.Names,
		Help:  helpCommand.Help,
		Run: func(args []string) {
			help.ShowCommandsHelp(append(commands, helpCommand), args)
		},
	})

	if len(os.Args) > 1 {
		cmdline.RunSubCommand(commandsWithHelp, os.Args[1:])
	} else {
		help.ShowCommandsHelp(commandsWithHelp, []string{})
		util.Exit(false, 1)
	}
}
