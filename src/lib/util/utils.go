package util

import (
	"bufio"
	"fmt"
	"os"
)

func ExitError(err error, pause bool) {
	fmt.Fprintln(os.Stderr, "ERR: "+err.Error())
	Exit(pause, 1)
}

func Exit(pause bool, exitCode int) {
	if pause {
		fmt.Fprintf(os.Stdout, "press enter to continue...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}

	os.Exit(exitCode)
}
