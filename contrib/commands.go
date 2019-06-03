package contrib

import (
	"errors"
	"fmt"
	"github.com/hellgate75/vmklet/commands"
	"github.com/hellgate75/vmklet/model"
)

func NewCommand(commandType model.CommandType) (*commands.ShellCommand, error) {
	switch commandType {
	case model.ShellCommand:
		return &commands.ShellCommand{
			Type: commandType,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Command type '%s' not provided!!", commandType.String()))
}
