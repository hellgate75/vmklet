package commands

import (
)

/*
* Http Command
* to enquire one or more url on remote/local host machine
* using multiple protocols
 */

//type ShellCommand struct {
//	Type		model.CommandType
//}
//
//func (Command *ShellCommand) Execute(stream model.CommandStream, commands ...string) error {
//	if len(commands) > 0 {
//		if len(commands) == 1 {
//			return  stream.Command(commands[0])
//		} else {
//			return  stream.Procedure(commands)
//		}
//	}
//	return  errors.New(errors.New(fmt.Sprintf("Shell has inappropriate number of commands '%d'!!", len(commands) )))
//}
//
//func (Command *ShellCommand) Parse(arguments ...string) (*model.Command, error) {
//	return &ShellCommand{Type: model.ShellCommand}, nil
//}
