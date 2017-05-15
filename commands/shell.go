package commands

import (
	"vmklet/model"
	"errors"
	"fmt"
)

/*
* Shell Command
* Used to execute shell commands on remote/local host machine
 */

type ShellCommand struct {
	Type		model.CommandType
}

func (Command *ShellCommand) Execute(stream model.CommandStream, commands ...string) error {
	if len(commands) > 0 {
		if len(commands) == 1 {
			return  stream.Command(commands[0])
		} else {
			return  stream.Procedure(commands)
		}
	}
	return  errors.New(errors.New(fmt.Sprintf("Shell has inappropriate number of commands '%d'!!", len(commands) )))
}

func (Command *ShellCommand) Parse(arguments ...string) (*model.Command, error) {
	return &ShellCommand{Type: model.ShellCommand}, nil
}
