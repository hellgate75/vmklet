package streams

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"time"
	"vmklet/model"
	"errors"
	"strconv"
	"vmklet/utils"
)

/*
* SSH Command stream
* SSH commands executor implements model.CommandStream
 */

type SSHCommandStream struct {
	Type					model.StreamType
	User         	string
	IP           	string
	Port         	int
	IdentityFile 	string
	Arguments    	[]string
}

func (ssh *SSHCommandStream) Command(cmd string) error {
	return  ssh.executeOne(cmd)
}

func (ssh *SSHCommandStream) Procedure(cmd ...string) error {
	return ssh.executeMany(cmd)
}

func (ssh *SSHCommandStream) Parse(arguments ...string) (*model.CommandStream, error) {
	var err error
	var user string = ""
	var key string = ""
	var ipAddress string = ""
	var port int = 0
	var arguments []string = make([]string, 0)
	var parsedCommand string = model.SSHStream.String() + " " + strings.Join(arguments, " ")
	if len(arguments) > 1 {
		for index,argument := range arguments {
			if index % 2 == 0 {
				//First token
				var key string = argument
				if index >= len(arguments-1) {
					return  nil, errors.New(errors.New(fmt.Sprintf("SSH Command '%s' has uncompleted argument '%s'!!", parsedCommand, key )))
				}
				var value string = arguments[index+1]
				if strings.Index(utils.CorrectInput(key), "user") >= 0 {
					user = strings.TrimSpace(value)
				} else if strings.Index(utils.CorrectInput(key), "identity") >= 0 {
					key = strings.TrimSpace(value)
				} else if strings.Index(utils.CorrectInput(key), "ipaddress") >= 0 {
					ipAddress = strings.TrimSpace(value)
				} else if strings.Index(utils.CorrectInput(key), "port") >= 0 {
					port, err = strconv.Atoi(strings.TrimSpace(value))
					if err != nil {
						return  nil, err
					}
				} else if strings.Index(utils.CorrectInput(key), "arguments") >= 0 {
					var args []string = make([]string, 0)
					for i:=index+1; i < len(arguments); i++ {
						args = append(args, strings.TrimSpace(arguments[i]))
					}
					arguments = args
					break
				}
			}
		}
		if user == "" || ipAddress == "" {
			return nil, errors.New(errors.New(fmt.Sprintf("SSH Command '%s' user and ip address are mandatory fields!!", parsedCommand)))
		}
		return  &SSHCommandStream{
			Type: model.SSHStream,
			User: user,
			IP: ipAddress,
			Port: port,
			IdentityFile: key,
			Arguments: arguments,
		}, nil
	}
	return nil, errors.New(errors.New(fmt.Sprintf("SSH Command has inappropriate number of arguments '%d'!!", len(arguments) )))
}

func (ssh *SSHCommandStream) executeOne(cmd ...string) *exec.Cmd {
	var args []string = make([]string, 0)
	if ssh.IdentityFile != "" {
		args = append(args, "-i", ssh.IdentityFile)
	}

	if len(ssh.Arguments) > 0 {
		args = append(args, ssh.Arguments...)
	}

	args = append(args, fmt.Sprintf("%s@%s:%d", ssh.User, ssh.IP, ssh.Port))

	args = append(args, cmd...)

	return exec.Command("ssh", args...)
}

func (ssh *SSHCommandStream) executeMany(commands ...string) ([]string, error) {
	var err error
	var out []string = make([]string, 0)
	var args []string = make([]string, 0)
	if ssh.IdentityFile != "" {
		args = append(args, "-i", ssh.IdentityFile)
	}

	if len(ssh.Arguments) > 0 {
		args = append(args, ssh.Arguments...)
	}

	args = append(args, fmt.Sprintf("%s@%s:%d", ssh.User, ssh.IP, ssh.Port))

	cmd := exec.Command("ssh", args...)
	reader, err := cmd.StdoutPipe()
	scanner := bufio.NewScanner(reader)
	defer reader.Close()
	err = scanner.Err()
	if err == nil {
		for _, command := range commands {
			cmd.Stdin = strings.NewReader(command)
			for scanner.Scan() {
				err = scanner.Err()
				if err != nil {
					break
				}
				out = append(out, string(scanner.Bytes()))
			}
		}
	}

	// Graceful ssh exit
	cmd.Stdin = strings.NewReader("exit")

	time.Sleep(2 * time.Second)

	if cmd.Process.Pid > 0 {
		// Forced ssh exit
		cmd.Process.Kill()
	}

	return out, err
}
