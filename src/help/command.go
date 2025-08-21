package help

import (
	"fmt"

	"github.com/reiyuchan/workshopdl/lib/cmdline"
)

func GetCommand() cmdline.Command {
	return cmdline.Command{
		Names: []string{"help", "--help", "-h", "/?"},
		Help: func() string {
			return getHelpMessage()
		},
		Run: func(args []string) {
			fmt.Println("Usage:")
			fmt.Println(getHelpMessage())
		},
	}
}

func ShowCommandsHelp(commands []cmdline.Command, args []string) {
	helpStr := ""
	if len(args) > 0 {
		command, err := cmdline.FindCommandByName(commands, args[0])
		if err == nil {
			if command.Help != nil {
				help := command.Help()
				if help != "" {
					helpStr += "\n" + help + "\n"
				}
			}
		}
	}

	if len(args) == 0 || helpStr == "" {
		for _, c := range commands {
			if c.Help != nil {
				help := c.Help()
				if help != "" {
					helpStr += "\n" + help + "\n"
				}
			}
		}
	}

	helpStr = "Usage:" + helpStr
	fmt.Print(helpStr)
}
