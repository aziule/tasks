package command

type Command interface {
	Execute() error
	LoadFlags(args []string) error
}
