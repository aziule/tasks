package command

type Command interface {
	GetName() string
	Execute() error
}