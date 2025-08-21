package cmdline

type Command struct {
	Names []string
	Help  func() string
	Run   func(args []string)
}
