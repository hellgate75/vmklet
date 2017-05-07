package model

type Storage interface {
	Save(file string) error
	Load(file string) error
}

type CommandStream interface {
	Command(cmd Command) error
	Procedure(cmd ...Command) error
}

type Command interface {
	Execute(arguments ...string) error
	Parse(commands ...string) error
}
