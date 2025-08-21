package get

import (
	"errors"

	"github.com/reiyuchan/workshopdl/lib/cmdline"
	"github.com/reiyuchan/workshopdl/lib/util"
)

func GetCommand() cmdline.Command {
	return cmdline.Command{
		Names: []string{"get"},
		Help: func() string {
			return getHelpMessage()
		},
		Run: execute,
	}
}

func execute(args []string) {
	if len(args) == 1 {
		err := WorkshopGetItem(args[0])
		if err != nil {
			util.ExitError(err, true)
		}
	} else {
		util.ExitError(errors.New("ERR: not enough arguements try using help"), false)
	}
}
