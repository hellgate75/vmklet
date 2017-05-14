package contrib

import (
	"vmklet/commands"
	"vmklet/model"
	"fmt"
	"errors"
)

func NewCommand(commandType model.CommandType) (*commands.ShellCommand, error) {
	switch commandType {
		case model.ShellCommand:
			return  &commands.ShellCommand{
				Type: commandType,
			}, nil
	}
	return  nil, errors.New(fmt.Sprintf("Command type '%s' not provided!!", commandType.String()))
}