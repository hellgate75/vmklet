package model

type Storage interface {
	Save(file string) error
	Load(file string) error
}

type CommandStream interface {
	Command(cmd string) error
	Procedure(cmd ...string) error
	Parse(arguments ...string) (CommandStream, error)
}

type Command interface {
	Execute(stream CommandStream, commands ...string) error
	Parse(arguments ...string) (Command, error)
}
