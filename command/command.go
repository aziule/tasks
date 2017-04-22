package command

type Command interface {
	GetName() string
	Execute(args []string) error
}
