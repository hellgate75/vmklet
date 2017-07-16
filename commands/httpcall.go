package commands

import (
	"fmt"
	"github.com/hellgate75/vmklet/model"
	"errors"
	"strings"
	"github.com/hellgate75/vmklet/utils"
)

/*
* Http Command
* to enquire one or more url on remote/local host machine
* using multiple protocols
*/

type URLCommand struct {
	Type		model.CommandType
	SSL			bool
	Format	string
	Timeout	int
	Code		int
}

func (Command *URLCommand) Execute(stream model.CommandStream, commands ...string) error {
	if len(commands) > 0 {
		if len(commands) == 1 {
		} else {
		}
	}
	return  errors.New(errors.New(fmt.Sprintf("Shell has inappropriate number of commands '%d'!!", len(commands) )))
}

func (Command *URLCommand) Parse(arguments ...string) (*model.Command, error) {
	var err error
	var ssl bool = false
	var format string = ""
	var timeout int = 200
	var expectedCode int = 200
	var parsedCommand string = model.URLCommand.String() + " " + strings.Join(arguments, " ")
	if len(arguments) > 1 {
		for index, argument := range arguments {
			if index % 2 == 0 {
				//First token
				var key string = argument
				if index >= len(arguments - 1) {
					return nil, errors.New(errors.New(fmt.Sprintf("URL Command '%s' has uncompleted argument '%s'!!", parsedCommand, key)))
				}
				var value string = arguments[index + 1]
				if strings.Index(utils.CorrectInput(key), "ssl") >= 0 {
					ssl = utils.GetBoolean(value)
				} else  if strings.Index(utils.CorrectInput(key), "format") >= 0 {
					format = utils.CorrectInput(value)
					if format != "http" {
						return nil, errors.New(errors.New(fmt.Sprintf("URL Command '%s' key '%s' not provided!!", parsedCommand, format)))
					}
				} else  if strings.Index(utils.CorrectInput(key), "timeout") >= 0 {
					timeout, err = utils.StringToInt(value)
					if err != nil {
						return nil, errors.New(errors.New(fmt.Sprintf("URL Command '%s' error parsing timeout argument '%s'!!", parsedCommand, value)))
					}
				} else  if strings.Index(utils.CorrectInput(key), "code") >= 0 {
					expectedCode, err = utils.StringToInt(value)
					if err != nil {
						return nil, errors.New(errors.New(fmt.Sprintf("URL Command '%s' error parsing return code argument '%s'!!", parsedCommand, value)))
					}
				}
			}
		}
	}
	if format == "" {
		return nil, errors.New(errors.New(fmt.Sprintf("URL Command '%s' fromat argument is mandatory!!", parsedCommand)))
	}
	return &URLCommand {
		Type: model.URLCommand,
		SSL: ssl,
		Format: format,
		Timeout: timeout,
		Code: expectedCode,
	}, nil

}
